package curl2go

import "testing"

func TestParseCommand(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "curl",
			args:    args{command: "curl -H 'Host: shopping.ele.me' -H 'x-cmd-v: 0%7C0' -H 'x-mini-wua: HHnB_ZmgCKqIxRZELq2FxrnpfOFytOtZR3iYqWQnOlKR%2BoXI4EBW7AlBxNdwHqdjm5rb0AAa6W4VbIuyNjEQhNHVoBtltdmQQ599uazXYN6qvvKMziHnOAvrBulML%2B%2BArUzdn' -H 'x-nq: WiFi' -H 'Accept: */*' -H 'x-location: 113.099319%2C28.242743' -H 'x-page-url: https%3A%2F%2Ftb.ele.me%2Fwow%2Falsc%2Fmod%2Fafda405d92b60a86b20538cc' -H 'User-Agent: MTOPSDK%2F1.9.3.48%20%28iOS%3B13.4.1%3BApple%3BiPhone12%2C1%29' -H 'x-features: 1035' -H 'Cookie: isg=BCUlFp9klnGLgfItj_828RNNPuVfYtn0_wXZticK4dxrPkWw77LpxLPczCRIJfGs; l=eBQ6hGcVOiJr4tWkXOfahurza77OSCdYYuPzaNbMiOCPOw165bbcWZy7OC8BC31Rh6EMR3SfuzDXBeYBmnA1mY_JCNOkkEHmn; tfstk=cZkNBAsUijcFEzw8IRwVhxtYpMXOZTDmesE_sjVnBJEJNueGMoeARRtMAH9AM; UTUSER=59447346; cna=RQbcF6DbdgsCAXjjEWXHpjNu; ut_ubt_ssid=5728eiz7ua0piu9au7pq55p4xx3z1spm_2020-09-07; xlly_s=1; DNT=0; x5check_ele=KsxPWTRBlvDotuQrJM4b5g%3D%3D; SID=BgAAAAADixgy7QEGHwBgO3kHl7Y1Tu-jydT6v6MoLqxMAofo7wq-PsNN; USERID=59447346; _tb_token_=535350ee155eb; cookie2=1cb1bf4e2461565234cd2687ffbea005; csg=c6255b88; munb=2204316740587; sgcookie=W100ZWUYGvtuPbdUXBUbmVnffs3lQn7EaF6lD9jpNobsFxmRKT0KCAt7Gx%2FWPO%2Fgdjdl3%2BeXQTPTS4BEVucnerej47hOyK71%2BDhrpygC40TBnS4%3D; t=0af58286bda80cb4a1e4ed850d2c62f4; unb=2204316740587; tzyy=50cadb2df19e17d3b36474f60a35d529' -H 'x-uid: 2204316740587' -H 'x-app-ver: 9.2.7' -H 'x-pv: 6.3' -H 'pageUrl: https%3A%2F%2Ftb.ele.me%2Fwow%2Falsc%2Fmod%2Fafda405d92b60a86b20538cc' -H 'f-refer: mtop' -H 'x-umt: LRpLosVLOjEphzV0ZG6KbpqtYbiVRSd0' -H 'x-sgext: JAGE458Fo2eIDfS4KB3eYQ%3D%3D' -H 'x-sign: izMRVI002xAAI90G9EOexsMzgdbvY90D0mK4Vim022q%2F0nm58VduoM%2BbrlSbRTWlNKkLGFRXj3SeWRkHjTqZdA80D6PdE90D3RPdA9' -H 'x-app-conf-v: 0' -H 'Accept-Language: zh-cn' -H 'X-Sufei-Token: eBQ6hGcVOiJr4qr3BOfChurza77TSCdbYuPzaNbMiOCPOafw5bbcWZy7OILeCn1Rh6A2R3SfuzDXBeYBmnxmnxv9CNOkkeHqn' -H 'x-appkey: 24894833' -H 'x-sid: 1cb1bf4e2461565234cd2687ffbea005' -H 'x-ttid: 201200%40eleme_iphone_9.2.7' -H 'x-t: 1599412301' -H 'x-utdid: X1UW9q%2FikbQDAFS3KSUA6ygZ' -H 'x-bx-version: 6.5.6' -H 'Content-Type: application/x-www-form-urlencoded; charset=utf-8' -H 'x-page-name: TBHomeViewController' -H 'x-ua: Rajax%2F1%20Apple%2FiPhone12%2C1%20iOS%2F13.4.1%20Eleme%2F9.2.7%20ID%2FDE4FE255-77D0-4F96-834F-2A8379CCB268%3B%20IsJailbroken%2F1%20ASI%2F3E6538A8-D8D0-44A1-8CC7-518C8867A0D7%20Mozilla%2F5.0%20%28iPhone%3B%20CPU%20iPhone%20OS%2013_4_1%20like%20Mac%20OS%20X%29%20AppleWebKit%2F605.1.15%20%28KHTML%2C%20like%20Gecko%29%20Mobile%2F15E148%20AliApp%28ELMC%2F9.2.7%29%20UT4Aplus%2F0.0.4%20WindVane%2F8.6.1%20828x1792' --data-binary \"data=%7B%22source%22%3A%22ALSC_MOD%22%2C%22pageId%22%3A%22314241%22%2C%22pvUuid%22%3A%22mAw1599412293844%22%2C%22locationInfos%22%3A%22%5B%7B%5C%22lng%5C%22%3A113.09931945800781%2C%5C%22lat%5C%22%3A28.24274253845215%2C%5C%22city%5C%22%3A%5C%22%5C%22%7D%5D%22%2C%22bizCode%22%3A%22KOUBEI%22%2C%22extData%22%3A%22%7B%5C%22passInfo%5C%22%3A%7B%7D%2C%5C%22modPreview%5C%22%3A%5C%22false%5C%22%2C%5C%22expectStatus%5C%22%3A%5C%22RUNNING%5C%22%7D%22%2C%22requestId%22%3A1599412301485%2C%22moduleRequestList%22%3A%22%5B%7B%5C%22moduleUuid%5C%22%3A%5C%227829694980%5C%22%2C%5C%22moduleName%5C%22%3A%5C%22emod-takeout-marketing-ten-billion-subsidy-shop-list%5C%22%2C%5C%22requestItems%5C%22%3A%5B%7B%5C%22resourceId%5C%22%3A%5C%222008142131000188751%5C%22%2C%5C%22distinct%5C%22%3Afalse%2C%5C%22extData%5C%22%3A%7B%5C%22minPerNum%5C%22%3A3%2C%5C%22perNum%5C%22%3A2%2C%5C%22moduleName%5C%22%3A%5C%22emod-takeout-marketing-ten-billion-subsidy-shop-list%5C%22%2C%5C%22moduleUuid%5C%22%3A%5C%227829694980%5C%22%7D%2C%5C%22startIndex%5C%22%3A20%2C%5C%22count%5C%22%3A20%7D%5D%7D%5D%22%7D\" --compressed 'https://shopping.ele.me/gw/mtop.alsc.activity.bpgw.eleme.batchquerymodule/1.0?rnd=43DF70B164A3ADF341B3FA1D97713B30'"},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCommand(tt.args.command)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
