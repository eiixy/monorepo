package helpers

import (
	"fmt"
	"testing"
	"zbanx-v2/internal/data/account/ent"
)

const originalJson = `{
        'channelType': 'Youtube',
        'quantity': 310000,
        'channelName': 'Andres Vidoza',
        'channelUrl': 'www.youtube.com/channel/UCC_NjLEb2Sley94py4vSYTA',
        'country': '加拿大',
        'language': 'en',
        'crawlerPartnerId': 20190135,
        'channelFrom': 4,
        'lastPostTime': 1683993600,
        'flow': None,
        'globalRank': None,
        'countryRank': None,
        'siteType': None,
        'siteTypeRank': None,
        'proportionTrend': None,
        'competitionUrl': None,
        'createTime': 1682535705.0,
        'remark': None,
        'issueModelTags': 'UGREEN 4k 60Hz,UGREEN USB C Hub,Plugable 5-in-1,Plugable USB C Hubs,Anker PowerExpand,Anker 551 USB-C Hub,Anker 7-in-1,Anker USB 3.0 Hub,EZQuest USB-C Multimedia,Dell 7-in-1 USB-C,Dell DA310,Dell DA300,HIEARCOOL USB C Hub,Hiearcool 8-in-1,Microsoft USB C Travel Hub,Sabrent 10-Port,SABRENT 4 Port,Sabrent 16 port,Belkin USB C Hub,Kingston Nucleum,ACASIS 10-in-1 USB-C Hub,Elecife 12-in-1',
        'issueBrandTags': 'UGREEN,Plugable,Anker,EZQuest,Dell,Hiearcool,Microsoft,SABRENT,Belkin,Kingston,ACASIS,Elecife',
        'issueProductLineTags': '拓展坞',
        'origChannelUrl': None,
        'mAvgNum': 6,
        'median': 187718,
        'avgViewNum': 243381.0,
        'avgLikeViewRate': 0.0237,
        'avgEngagementRate': 0.0249,
        'avgComment': 288.0,
        'avgLike': 5772.0,
        'videoCount': 0,
        'watchCount': 40572423,
        'partnerVideoList': [],
        'crawlerUserId': 58400219,
        'secUid': 'UCC_NjLEb2Sley94py4vSYTA',
        'sixVideoNums': 31,
        'videoNums': 153,
        'needFilter': 1,
        'playSum': 43897245,
        'userDefinedId': None,
        'crawlerTaskIds': ['a9db915b4a524c979d5cbb30bd648f9a'],
        'crawlerUpdateTime': 1684234023,
        'crawlerTaskDetails': [{
                'taskId': ['a9db915b4a524c979d5cbb30bd648f9a'],
                'recommendCount': 5,
                'recommendDetail': '作者被搜索词推荐统计:\nPlugable 5-in-1 (x1)\nDell 7-in-1 USB-C (x2)\nDell (x2)\n作者发帖被搜索词推荐统计:\n链接:https://www.youtube.com/watch?v=4jC8CeZmyUk 时间:2021-12-05 观看:276893 (x1) title:M1 MacBook Air vs 14” MacBook Pro M1 Pro | Programming & General Use\n链接:https://www.youtube.com/watch?v=jezxearJGP0 时间:2023-01-08 观看:296786 (x1) title:Real BUDGET Tech Gadgets UNDER $75\n链接:https://www.youtube.com/watch?v=OBb-DmJN9_U 时间:2023-04-09 观看:148458 (x1) title:Real BUDGET Student Accessories Worth The Money!\n链接:https://www.youtube.com/watch?v=exq5fQGFzjo 时间:2022-10-23 观看:120528 (x2) title:Back To School With the Dell XPS 13 Plus Review | M2 MacBook Air Alternative ?',
                'inKeywordCount': 0,
                'inKeywordDetail': [],
                'inKeywordReverseCount': 17,
                'inKeywordReverseDetail': [{
                        'url': 'https://www.youtube.com/watch?v=FoIuX0riX-U',
                        'text': '反链包含Anker'
                }, {
                        'url': 'https://www.youtube.com/watch?v=lfl2RcvliMY',
                        'text': '反链包含Kingston'
                }, {
                        'url': 'https://www.youtube.com/watch?v=QEGZPUoi5fA',
                        'text': '反链包含Belkin'
                }, {
                        'url': 'https://www.youtube.com/watch?v=xECTAuYuRwQ',
                        'text': '反链包含Dell'
                }, {
                        'url': 'https://www.youtube.com/watch?v=rojGkxDcACY',
                        'text': '反链包含Microsoft'
                }]
        }],
        'headPortrait': None,
        'relatedLinks': [{
                'originalLink': 'https://twitter.com/andres_vidoza',
                'afterLink': 'https://twitter.com/andres_vidoza',
                'secUid': None,
                'effective': True
        }, {
                'originalLink': 'https://www.tiktok.com/@andresvidoza',
                'afterLink': 'https://www.tiktok.com/@andresvidoza',
                'secUid': None,
                'effective': True
        }, {
                'originalLink': 'www.canvogh.com',
                'afterLink': 'https://canvogh.com/',
                'secUid': None,
                'effective': True
        }],
        'countryCode': 'CA',
        'avgViewRate': 0.7954,
        'avgViewGrade': 0.7954
}`

func TestGzip(t *testing.T) {
	s := originalJson
	encode, err := GzipEncode([]byte(s))
	if err != nil {
		return
	}
	fmt.Println(len(s), len(encode), float64(len(encode))/float64(len(s)))
	decode, err := GzipDecode(encode)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decode))
}

func TestMarshalJson(t *testing.T) {
	var m = ent.Account{
		ID:       1,
		UID:      "test_uid",
		Username: "test_usernmae",
	}
	andGzip, err := MarshalJsonAndGzip(m)
	if err != nil {
		panic(andGzip)
	}
	fmt.Println("MarshalJsonAndGzip", string(andGzip))

	var m1 = ent.Account{}
	err = UnmarshalDataFromJsonWithGzip(andGzip, &m1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("UnmarshalDataFromJsonWithGzip: %+v\r\n", m1)
}
