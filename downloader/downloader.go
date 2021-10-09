// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-10-08 15:39
// version: 1.0.0
// desc   :

package downloader

import (
	"e-learn/logger"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
	"github.com/yhyzgn/goat/file"
	"github.com/yhyzgn/golus"
	"io"
	"math"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Resource struct {
	path     string
	filename string
	url      string
}

type Downloader struct {
	wg         *sync.WaitGroup
	pool       chan *Resource
	Concurrent int
	HttpClient http.Client
	TargetDir  string
	Resources  []Resource
}

func New(targetDir string) *Downloader {
	return &Downloader{
		wg:         &sync.WaitGroup{},
		Concurrent: runtime.NumCPU() / 2,
		TargetDir:  targetDir,
	}
}

func (d *Downloader) Append(path, filename, url string) {
	d.Resources = append(d.Resources, Resource{
		path:     path,
		filename: filename,
		url:      url,
	})
}

func (d *Downloader) download(resource Resource, progress *mpb.Progress) error {
	defer d.wg.Done()

	d.pool <- &resource

	dir := path.Join(d.TargetDir, resource.path)
	if !file.Exists(dir) {
		err := os.MkdirAll(dir, os.ModeDir)
		if err != nil {
			return err
		}
	}
	finalFilename := path.Join(dir, resource.filename)

	// 文件不存在时才下载
	if !file.Exists(finalFilename) {
		logger.InfoF("开始下载文件【{}】...", finalFilename)

		// 创建临时文件
		tempFilename := finalFilename + ".tmp"
		target, err := os.Create(tempFilename)
		if nil != err {
			return err
		}

		req, err := http.NewRequest(http.MethodGet, resource.url, nil)
		if nil != err {
			_ = target.Close()
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if nil != err {
			_ = target.Close()
			return err
		}

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				panic(err)
			}
		}(resp.Body)

		proxyReader := resp.Body
		if nil != progress {
			// 获取到文件大小
			fileSize, _ := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
			labelRune := []rune(finalFilename)
			label := string(labelRune[int(math.Max(float64(len(labelRune)-32), 0)):])
			bar := progress.Add(
				fileSize,
				mpb.NewBarFiller(mpb.BarStyle().Lbound("[").Filler("=").Tip(">").Padding("_").Rbound("]")),
				mpb.BarFillerClearOnComplete(),
				mpb.PrependDecorators(
					decor.Name(label, decor.WC{W: len(label) + 1, C: decor.DidentRight}),
					decor.CountersKibiByte("% .2f / % .2f", decor.WC{W: 32}),
					decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_MMSS, 0, decor.WCSyncWidth), ""),
				),
				mpb.AppendDecorators(
					decor.OnComplete(decor.NewPercentage("%.2f", decor.WC{W: 7}), "  "+golus.New().FontColor(golus.FontGreen).FontStyle(golus.StyleBold).Apply("Download Finished")),
				),
			)
			proxyReader = bar.ProxyReader(proxyReader)
		}

		defer func(proxyReader io.ReadCloser) {
			err := proxyReader.Close()
			if err != nil {
				panic(err)
			}
		}(proxyReader)

		// 将下载的文件流拷贝到临时文件
		//_, err = io.Copy(target, proxyReader)
		_, err = io.CopyBuffer(target, proxyReader, make([]byte, 32*1024*1024))
		if nil != err {
			_ = target.Close()
			return err
		}
		_ = target.Close()

		logger.InfoF("文件【{}】下载完成", finalFilename)

		// 修改临时文件为最终文件
		err = os.Rename(tempFilename, finalFilename)
		if nil != err {
			return err
		}
	}

	<-d.pool

	return nil
}

func (d *Downloader) Start() error {
	d.pool = make(chan *Resource, d.Concurrent)
	logger.InfoF("开始下载，当前并发数：{}", d.Concurrent)

	progress := mpb.New(
		mpb.WithWaitGroup(d.wg),
		mpb.WithWidth(60),
		mpb.WithRefreshRate(300*time.Millisecond),
	)

	for _, resource := range d.Resources {
		d.wg.Add(1)
		go func(res Resource) {
			err := d.download(res, progress)
			if err != nil {
				return
			}
		}(resource)
	}

	progress.Wait()
	d.wg.Wait()
	return nil
}
