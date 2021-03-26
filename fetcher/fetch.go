package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const tmpCookie = "sid=b6ac8ebe-f1d7-45de-8071-460194548fd4; ec=g5s2spdj-1616480328670-d49e7f968231c1248756317; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1616480336; FSSBBIl1UgzbN7NO=5OlFjG_.fxShf3yo9inUrzz7ACHfPwtA6qpHzla3aAwR6z1hDQfb6OXVwrdqK2vml0aQ5Y1CEQ83srv.hHQZT4G; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1616726093; _efmdata=krmqwcCi5F6JC0fs0IU3OI83eFfH6kQ%2BeeLWsRiRLHuQAwAXcfd%2Bst7vGR6sJpCvYYnN3UgKTXkTSv2LTRhSPonODtsByjKNsV0ebpivE3I%3D; _exid=8BbiQdyxcV7Juch%2BmEHQL0w1fXGJYLS%2F0et%2Bhez%2FGfHsaKtHI9XQv4t5JsF9EYVgAmV3U2%2FqGT2HF3KbqztUpg%3D%3D; FSSBBIl1UgzbN7NP=53mTEsCrw1laqqqmgMDW4iGElonpal6S5FBNaWQ01GKdNPskcHDCAVInOTU0wxiM40hBD1IYsF8zQz0DjTXD8SvPDFxIYkX_dJKwSJV8jR9zl3eaZn2erKeyD7U_DVRsOwiBHRYE5RYIQqRUi6Ut_EG3dzsNJP9gpOHZIuI_WZnVghe9Qx_wD6Lx0PQX3Wrin4eVFkeXBAp30ICgjIf7NstlsZATkI_xcjsWNa9dT9sqBRP._qifgA_5CZoUtFnSGsAEsTod_lGmP6lhMbE63j8.AyzXylvO2BxGkt2wDQPSXwM7hVYg6Y7VteQKvD5XU0"

func Fetch(url string) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)

	req.Header.Add("cookie", tmpCookie)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("err status code %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
