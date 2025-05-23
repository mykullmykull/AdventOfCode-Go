package day

// ------------------------------ inputs ------------------------------------//
var testInput = []string{
	"root: pppw + sjmn",
	"dbpl: 5",
	"cczh: sllz + lgvd",
	"zczc: 2",
	"ptdq: humn - dvpt",
	"dvpt: 3",
	"lfqf: 4",
	"humn: 5",
	"ljgn: 2",
	"sjmn: drzm * dbpl",
	"sllz: 4",
	"pppw: cczh / lfqf",
	"lgvd: ljgn * ptdq",
	"drzm: hmdt - zczc",
	"hmdt: 32",
}
var expected = 301

var realInput = []string{
	"rpws: 4",
	"cnnj: ljdg * gmcc",
	"npss: 5",
	"gfds: fgmh * scwv",
	"brhs: nbfp * fcvr",
	"lzfd: 17",
	"fbhc: sflc + vnnj",
	"jcmv: 14",
	"pqrc: cngs + ddtq",
	"lgss: wpwj * pdjl",
	"gmch: 2",
	"mzbw: 5",
	"tsbt: rbvs + htpp",
	"wsfw: 10",
	"rbpp: 7",
	"nhcw: jnrj + sfpn",
	"fvhp: hnzd + ngft",
	"hdnl: 2",
	"qwpg: 5",
	"gmfj: wcff + wltd",
	"ssvh: dlct * jwpr",
	"djgn: 9",
	"llsw: 2",
	"qcht: 1",
	"cfqs: slws - fjzn",
	"cshb: gpfj + dbgn",
	"nnlh: lrjw * cnqn",
	"bqlv: rddq * glll",
	"chzs: hzdn + njmh",
	"rhzn: 5",
	"ffvp: 3",
	"zqls: 5",
	"rtfm: mcqd - srgg",
	"lzdd: 3",
	"nsdq: jsmw * tmvv",
	"vclq: jhch * swfj",
	"rhqz: 2",
	"tvcp: bvct + wwlp",
	"rhbl: wblz + svhn",
	"ncrr: qtlz + npdq",
	"ttsh: 3",
	"wjll: 4",
	"czgv: 2",
	"zbnp: 3",
	"prwl: mvcb - ldcs",
	"pmjn: pfrq * ldqj",
	"mvvb: jzmv + pwrw",
	"bgrb: szgz * tfjg",
	"djhz: 17",
	"scbg: jnhz / bzzf",
	"bpfh: 14",
	"nzhs: jwhn + jrlf",
	"gbgp: 10",
	"mclp: 1",
	"tfsh: 3",
	"vpwh: vpsr - pmjn",
	"dcbt: djgq * pcvl",
	"rfts: wmsm * zrvj",
	"rvbj: 20",
	"hwnj: zjmr * svwt",
	"vrgz: 3",
	"dczh: tnnv * qjmp",
	"rwsr: 5",
	"zjpf: 14",
	"rbgr: frpn + mvlq",
	"tvjv: tjzm + lsjq",
	"plsw: 3",
	"swjg: 16",
	"hfbq: wsln * tsss",
	"cdnm: dgzv + psnn",
	"fpsm: wrjw + hptw",
	"zflm: 5",
	"lnvd: jmnm + vcmr",
	"tbdd: sjzb + wbtm",
	"grqf: 2",
	"zggt: 9",
	"pjql: 5",
	"hjlw: fvnb / tsnd",
	"hhzp: 1",
	"pfht: 5",
	"ggqd: nggg / rrbn",
	"mqhh: 8",
	"svrj: 5",
	"hrnb: blml * bjpb",
	"dnhn: qrcb + mlng",
	"mhbb: ftjl + gjgg",
	"fjqm: thwh + nzzj",
	"sqfv: 5",
	"ldqj: lmsm + mzdj",
	"wfmj: 2",
	"lrrf: 4",
	"bqlg: zwds * lbzp",
	"wbtm: 4",
	"vsnl: glvr + ztmc",
	"jvps: 2",
	"slvj: zlmg + lbvs",
	"bjnn: zmgr * qhtc",
	"zvrb: 3",
	"ltwz: vztd * hdph",
	"rrdj: jqtc / hblw",
	"tctd: 2",
	"bnlm: qsgm + drfl",
	"ptqv: 3",
	"twpr: 2",
	"dljc: fdvh + wztz",
	"mhcp: bcbw - tlst",
	"qbmc: 2",
	"vlvc: lzzn + hflp",
	"mcqd: rfqt * lgzf",
	"nhsq: qtjw * pftc",
	"njlj: 3",
	"vvfg: nzhs * szwb",
	"dhpv: fhbg * dqgj",
	"dhqb: 5",
	"ldct: tsfc * bwmp",
	"tvbs: 4",
	"qlgm: 13",
	"prlz: gjtq * gcln",
	"jsdt: 12",
	"mltf: cfqs + mpch",
	"bwvw: 11",
	"cggw: 3",
	"qwcr: qscp / jdsv",
	"dqlw: znlm + scqb",
	"gmmz: dhvm + mclp",
	"qggr: nhsq + cpjr",
	"sbvt: zpmp * wdzb",
	"jdsv: 3",
	"lwfz: 2",
	"zgnl: zmvg + mhjb",
	"gvrs: 2",
	"sbsd: 1",
	"dgsm: vtwn - hjbd",
	"ndnp: 7",
	"nwhh: gprr + sbpt",
	"mstb: trtt - cpsb",
	"qhgg: 7",
	"bsvd: 2",
	"gjtq: 2",
	"vslc: hfbq + nqwc",
	"rnsn: 3",
	"hcwr: 4",
	"mhpf: 15",
	"hnzt: bzzs * nghf",
	"tjfm: lqgt + cmrz",
	"mcbz: 4",
	"mqhc: 3",
	"jsbq: vdht + jqmn",
	"wcff: 14",
	"zqdw: 17",
	"zdqg: mnds * rczw",
	"rssg: cjsv + lsfr",
	"bcbw: hpbb * lsqw",
	"htwz: lljf * pblp",
	"zdwj: vcdw + cfmm",
	"vtvs: 3",
	"qrll: dqlw * mbrd",
	"qzqh: 3",
	"spjh: 2",
	"dgtd: ztbg + msvr",
	"tfjg: mwjw * snzv",
	"gdtv: 3",
	"tddj: 4",
	"mgcp: 4",
	"nqcc: tlnd * wsjm",
	"zpdp: wtdz + bfjd",
	"cnqh: tlrh * lzhh",
	"ffvt: jrln * vdgg",
	"tfbg: 9",
	"bfjg: 2",
	"zwgm: 5",
	"vmvf: bbhw + gqsn",
	"vphf: 6",
	"blmn: ptrj * vbdh",
	"znlm: 1",
	"pbdw: sgww * hjwt",
	"drtj: 2",
	"qdgd: vgsj * mplj",
	"hphh: 17",
	"mwcg: 2",
	"tvmd: lrps * stnz",
	"wbcg: crzr * cjgd",
	"mspm: lgnd + nwhh",
	"gbbl: llsr * qvzv",
	"jlrs: 4",
	"ncmh: ggqd * jbzr",
	"stnz: 2",
	"cphz: 3",
	"vrwv: 2",
	"rtth: 2",
	"clnd: 3",
	"gqwn: lzmf * gpvn",
	"brbr: 1",
	"dhdq: 2",
	"ptrj: sjhq + pfht",
	"frls: 4",
	"jscp: 4",
	"dpth: wwdg + clww",
	"qgnm: nhcw / vqtw",
	"fnml: sbsd + qrtz",
	"vmbm: 2",
	"pcjz: 2",
	"jljw: 8",
	"bntv: 7",
	"nwlg: 2",
	"mlng: dnqd * pmgn",
	"fmnb: 15",
	"dqwh: 2",
	"gmbg: 2",
	"cctn: 11",
	"hfnv: 5",
	"wcds: 6",
	"jnsd: fsnc * ncgd",
	"bjng: 1",
	"rtqp: 2",
	"wffw: 5",
	"zpdg: 3",
	"hblw: 10",
	"mmnb: 17",
	"vbnc: 7",
	"csnm: gcvq * hrnb",
	"zlwm: bgrb + mltf",
	"rcrt: jlbb + rbtc",
	"sjsr: tqmm + nqjg",
	"nzzj: 2",
	"vzbq: nbcr / dswn",
	"sdrn: wpfs + pqlf",
	"rcst: dlsd * bhqr",
	"wmlz: jfjf * tsbt",
	"zqrj: tdfp + gjtr",
	"fcvr: 3",
	"mqrj: mfgg * jzjc",
	"ggjg: fnml + wjqf",
	"swtm: 13",
	"scms: qpfq * zwth",
	"lbgh: zgqh + lhzq",
	"jgvf: cldg + pgwq",
	"cslv: zntp * fwhf",
	"ntvl: 3",
	"dbch: 14",
	"mdpn: mdnc + hcdl",
	"tvtm: 4",
	"bqrw: qmfp * wfmj",
	"zwwq: jswh + dfmj",
	"jzgf: 4",
	"qdqq: mhgh * vjwj",
	"swpc: vsrr * zrcq",
	"hrjn: zgdt + vzzn",
	"qscp: sbvt - gndq",
	"ddrp: grvv - spdd",
	"mplj: 3",
	"dcll: lqqb + zszb",
	"sdbg: vqpg * lqgr",
	"zlpf: 2",
	"slrm: pzcc + hrvd",
	"bdvp: 5",
	"sccb: ftgs * jshq",
	"qtwl: 5",
	"gfzj: 4",
	"stnv: tswp + sdhd",
	"gtqd: 7",
	"lbrr: bbfd + mmnb",
	"nbrb: bnqj + nplr",
	"mlzt: nbrb * vmlz",
	"gmcc: 5",
	"sshr: 3",
	"bfzm: humn - nsqd",
	"hvmr: qtqw * gbpt",
	"fhbg: 5",
	"njmh: hmwf - zzqn",
	"pwrw: 3",
	"gztd: pvpl * jnwp",
	"vjch: 7",
	"prwd: tfbg - dvhn",
	"srdd: 1",
	"njvl: sjtl + qzng",
	"pdtb: 3",
	"wwtm: 16",
	"tqdj: rlql * vhrh",
	"cgtv: 5",
	"hmhf: gjpm * cwzj",
	"lgzf: qsfh * nqrv",
	"wcft: 4",
	"zsfw: 3",
	"hgfd: mjdd * vslc",
	"gjdl: zbnp * vshp",
	"jhch: gldz + hvgv",
	"wqvc: 14",
	"fqjh: 2",
	"spjf: 20",
	"hwff: zhpb + cgwc",
	"vprg: qzdn * mcbz",
	"vfqq: 6",
	"vnhn: 7",
	"hjbd: sfjp * nnlh",
	"pbgd: 1",
	"tdzq: hrjn - ddnn",
	"sqqn: phwn * trsc",
	"mvbg: 5",
	"cnht: 1",
	"cznv: bswd * rlfq",
	"wtbn: hgsf + bnls",
	"lszb: lwnf / znqt",
	"hflp: rfwb + vggg",
	"lsvb: 6",
	"pjbj: swtm * npsq",
	"hmvb: vmts * tqzm",
	"cwzj: 2",
	"bdlp: fbzf / rhzn",
	"rmnw: bbsz * mvvb",
	"fhns: 3",
	"qbzz: dzhc / zdww",
	"qdcf: wbcs * jmzd",
	"zzqn: vlzm / gqmm",
	"vffd: 9",
	"ttrz: 2",
	"nggg: grgc * lzlv",
	"wsln: 3",
	"thjd: vrqq + rhnr",
	"bmzn: pnmf - cdrs",
	"csdv: 2",
	"cjgd: 2",
	"blwh: lpqz / mmwh",
	"dbhd: whdl * tptq",
	"hczs: 3",
	"lmjr: 15",
	"rwqm: 5",
	"nwgr: jqgt * npgn",
	"cmth: 4",
	"fwbt: jlhp + ntzj",
	"vlbv: 10",
	"rqdn: dcbt / jzlm",
	"ffqn: bgzf + rcfn",
	"flbh: fpzj * fqmt",
	"fpqf: 2",
	"ptwr: glhq * mttf",
	"rdpq: 3",
	"mhjb: cznv * bsmj",
	"tzdl: 8",
	"phqr: 2",
	"wtgr: smdn + mrjz",
	"wrlb: znrh + lwwj",
	"jthd: tvjv + swpc",
	"jbcp: 5",
	"rhff: 4",
	"nzrt: jglh * jttc",
	"fpml: rlln * csnn",
	"qpjd: 3",
	"htpp: 3",
	"fqnl: 2",
	"mfdp: 6",
	"lrps: 4",
	"mwvj: tshz / lgrl",
	"svwt: gdrh / nphw",
	"rtmp: nqhn * zprq",
	"rwfw: twjf + tgrz",
	"ntgl: 10",
	"bzln: 3",
	"cbvv: 5",
	"grst: 3",
	"cdrs: mqsz + qbgf",
	"mnds: 5",
	"qjgd: 2",
	"lwbh: 13",
	"vdjh: sptz + rwfv",
	"nwnt: 2",
	"fnld: 2",
	"bzzs: 14",
	"zlmg: zlwm * dnmm",
	"hbvm: cfvj / qpjr",
	"dgns: 4",
	"cdgv: fddp + vzpl",
	"cpsb: 6",
	"tdvn: qfdv * lzfd",
	"shvl: glgv + cdgd",
	"cwnw: nbjm + prcm",
	"vjvd: sqgv + vjsl",
	"rrbn: 3",
	"tnrq: pzjj + zlpf",
	"llmr: fznq + dddr",
	"tzgm: wptd * wmrd",
	"fgjs: mwvj + nflc",
	"hrvd: zfjf + bfrr",
	"qlgh: nzrt / fjjn",
	"vdgg: 2",
	"blqz: hmvb - sgjd",
	"jpdg: 5",
	"jrlf: jjrm * rwfw",
	"lzzn: nfdb * vcgj",
	"bdjq: lcwp * jjfw",
	"bpwg: ssvh + rsqf",
	"nlgj: zlch * grcp",
	"lzqh: sfzf + zcjr",
	"cgdh: mhcp * vbtz",
	"vmvg: 3",
	"dltt: 4",
	"hgsf: cwns * bdjt",
	"dncc: 2",
	"gjgg: 1",
	"wsnm: dzjj + mwqj",
	"hlvn: 2",
	"lntg: 2",
	"twwc: 2",
	"cmrz: 7",
	"ctqd: 2",
	"jmrv: cgtv * rwqm",
	"sflc: 7",
	"msfm: 5",
	"cqjg: cnfj + qprc",
	"nhsv: sfvm + cpnl",
	"hptw: 4",
	"rfjd: 4",
	"jlmt: qdcd * bffz",
	"wpfc: fjqm * vqjh",
	"chgc: 13",
	"wptd: 3",
	"znrh: vsll / vfqq",
	"fbvj: 5",
	"rgvg: 3",
	"glvr: 19",
	"vgnc: 5",
	"zzdh: 5",
	"gdsl: drjw * qpch",
	"zlch: 2",
	"qwwh: sfrj * hhcs",
	"vbgd: 7",
	"nqjg: hmvc / smsg",
	"sppw: 5",
	"vgfb: 19",
	"lwmn: zbgb + svcv",
	"jrln: 20",
	"lwft: 10",
	"tmjw: 2",
	"htlw: whhm + jsdt",
	"hcdt: zszg - sprn",
	"lgnd: 2",
	"lqjh: wcnc + mpjh",
	"qprc: qwwh - rrcd",
	"gpvn: 9",
	"tsss: 9",
	"jbqb: 5",
	"lwqv: bjfv * cnbj",
	"fcbq: jzmf * hhpt",
	"vhps: bptz + pzjn",
	"chlr: 3",
	"zglw: 4",
	"vrgr: 2",
	"rgpp: wthc * rtbc",
	"bvtq: blmn + wgmb",
	"jhbw: 2",
	"gldz: zwzq + gzhp",
	"bzzf: 2",
	"wpvd: 5",
	"tlcq: lmbj + qszh",
	"fhjr: 4",
	"lszr: 2",
	"cnhr: 7",
	"jdbd: ddmh * dqtt",
	"pzcc: jmrl + lrmd",
	"ntph: 4",
	"sprn: 1",
	"bffz: 10",
	"pbtr: tvgv - qmdh",
	"sblz: 5",
	"bfng: 3",
	"zthj: mvsj + mdpn",
	"qmbd: 5",
	"jnls: bhzp * hlvn",
	"qfvh: fdms - jznt",
	"csvc: 2",
	"rsfg: lggb - lwgv",
	"nnmc: lzqh - qgws",
	"frss: 5",
	"sfhw: 4",
	"cgwc: vrgr * wffs",
	"pdsg: mpjz / dncc",
	"wzvr: 3",
	"zhgs: 2",
	"jzlh: 2",
	"zdqc: 4",
	"wjfb: svrm * nthz",
	"tpjd: 3",
	"ttpf: 4",
	"hmwf: zfvj * qtwl",
	"bpvd: 2",
	"jswh: rnwb * vbgd",
	"hgzt: 15",
	"fcmg: 2",
	"rczw: mqsv + mlws",
	"bgms: 11",
	"swqz: 5",
	"mwwz: 11",
	"qdsb: hrdh * jnbr",
	"wbpr: zscj * vwws",
	"rdjp: 5",
	"dgzv: vdjh + lhhr",
	"mjhh: wllr + pnld",
	"mpjp: 3",
	"cqhf: 5",
	"fgmh: 4",
	"wfqn: 4",
	"qvzv: sfzg * tvjp",
	"ftgs: ptwr - gncj",
	"glph: jhws * fwnz",
	"rbtc: lszr * lcbc",
	"rtsf: qpjl + vmvf",
	"dtsn: 11",
	"hrjd: zwwq + wbqt",
	"grdt: 3",
	"npgn: 2",
	"vtgv: ffrf * hwbg",
	"mzjh: 5",
	"sjhq: 2",
	"bgpn: 13",
	"cnfj: 3",
	"pbpj: mhcw + dqfb",
	"bjpb: 2",
	"gjpm: cmhq * mfcc",
	"bjfv: 2",
	"rnqb: 4",
	"jglh: sqfv * mvpp",
	"mtnm: vrwv * gfhl",
	"qgfj: qcht + vgqh",
	"bmvz: hrjd / bqsc",
	"grcp: bgdg * gpht",
	"fjzn: 1",
	"dzzc: mjhh / zhgs",
	"lsjq: dhzq * jnmq",
	"vfmd: 14",
	"tscs: sttl + jlmt",
	"vjdw: mtbl * dlwp",
	"sdwd: 4",
	"zbzp: 2",
	"mdvr: zgtd * qbmc",
	"phwn: twps - frls",
	"qzng: htlw * zpdg",
	"pdrz: npcz / gwbf",
	"zcjr: hlwc + mtzq",
	"sscn: 2",
	"rqjg: jljw * vsnl",
	"nbjz: 14",
	"sdhd: cdnm / fbvj",
	"zfcg: qhwr + cnqh",
	"cldg: pdtb * fvcw",
	"qvvp: 5",
	"pzlc: hlhl + jvdv",
	"zjlw: 3",
	"bpmg: lmfz * wjcp",
	"qcnw: btcl + ggrc",
	"srpl: 2",
	"lrmd: 19",
	"vqjh: 5",
	"rfrj: qhbc + htqf",
	"ntzj: jdgv + rqdn",
	"jjgj: vgrb * qbnf",
	"gncj: 9",
	"cqlj: vtvs * gfcs",
	"wcdz: 3",
	"qmfp: 8",
	"blff: cvmd * tzvp",
	"lwwj: jbtc - hcss",
	"pcsr: glnr * sscn",
	"tvss: 5",
	"dflq: zglh * nqzl",
	"bqrt: zbfl + gqwn",
	"czcd: bnlm + dfhj",
	"qltg: 2",
	"jzmf: 2",
	"nlpd: 5",
	"shvd: fpmh * gmgc",
	"mfqf: nggv + hszd",
	"ztpp: cnht + qgbj",
	"sgww: plcb * rdjp",
	"wltp: 5",
	"fvcw: lcph * jzrp",
	"wwdg: nsns + hhpb",
	"hcqg: 2",
	"pblp: 3",
	"qgvg: 5",
	"lzmj: 15",
	"bpbw: rpdn - dzrg",
	"gctm: whhd * hqpm",
	"fddp: 6",
	"rmcd: svsq + vmgp",
	"vvwv: 2",
	"cwsq: 2",
	"ltpw: lbqv * rtqp",
	"mrwf: 11",
	"vzvh: vvhc / dqnv",
	"jnwp: 16",
	"crzr: 3",
	"vtls: 8",
	"crgg: ntgn + gsgv",
	"vdqr: 8",
	"rsqf: hfvm * hwnj",
	"zfjf: vbgq * jwpp",
	"lhql: 5",
	"mfgg: 2",
	"gbcg: 2",
	"vsvt: 2",
	"vmjs: gldf + hmgr",
	"zcbg: 2",
	"szwb: 2",
	"fscn: tbwm + jnls",
	"hjqv: rnfc * sdnq",
	"zmbh: vnhn * rssg",
	"lbtr: 2",
	"zctc: 18",
	"mttn: qgdc * zpds",
	"hvgf: wsnm + snft",
	"hsvv: 4",
	"dqfb: wqvc + zgzt",
	"hrdv: 4",
	"rspj: mqhh + ptvw",
	"htlt: 3",
	"jdgv: cwdg * gwmj",
	"ndqv: 3",
	"fdgq: 2",
	"ldqh: 3",
	"gzhp: dtzs * gbcg",
	"fcnt: 9",
	"qwbn: llfz * vfsl",
	"sqqg: 14",
	"fvzh: 19",
	"gcln: vtls * gfhv",
	"tdfp: gmll / vftl",
	"gbzv: jcmv + cwrf",
	"ddtq: fcnt + sqwl",
	"vdht: 9",
	"nsqd: lqvl / trjw",
	"rnsh: 2",
	"dbgn: 3",
	"jvdv: gtwz * gtcd",
	"dhmm: 5",
	"lcjf: gvrs * lmrz",
	"jbmr: 4",
	"msjj: 4",
	"hhpc: 18",
	"jbsd: cwsq * rdps",
	"vgsj: 2",
	"srhc: zcbg + lmnc",
	"hmgr: 4",
	"pczj: vgwn + bcgj",
	"hhpb: ttpf * jlrn",
	"jvfh: mndb + jchl",
	"vnbl: 9",
	"vcmr: 4",
	"tlst: 4",
	"sdwt: 5",
	"dcrn: dpzw * rtsf",
	"lsbg: bsht * rfjd",
	"llsr: 5",
	"mlwn: 3",
	"ngsg: 3",
	"jhhv: 16",
	"ztmc: sdwd + bpws",
	"srwm: wnsb * tlbc",
	"rmbs: 5",
	"tvjp: 3",
	"qhwr: 2",
	"mtdt: bpqv * hwfd",
	"jmzd: 3",
	"pzjs: 3",
	"smvl: vpwh * ctmf",
	"gpsn: vglg + gfds",
	"rffd: bvtq + lsfc",
	"qjmp: 3",
	"trgn: 3",
	"zddf: lzdq + pztv",
	"swls: hlzl + tvcp",
	"tzhd: nhsv + rcst",
	"whhm: tfdm * fbwv",
	"vglg: 2",
	"tptq: bvhg + fsht",
	"spdd: 3",
	"bhqr: 2",
	"qdlg: wwgr - rpdq",
	"cgnf: 5",
	"vflz: 2",
	"rcll: mqph + mjrs",
	"plcb: ngnr + cqjh",
	"wdvw: vzpd + llrt",
	"vbtz: fzbj + zgnl",
	"dqcs: ltwz - blpd",
	"qwzs: 3",
	"fcpm: 3",
	"htrv: 7",
	"whqh: vfmd + sptt",
	"ldzz: lcjf - cghq",
	"sqgv: gdlp * rhqz",
	"vqpg: nqfw * zrzd",
	"wmsm: 2",
	"vsll: pwsf / sshr",
	"jfjf: tncp * wfqn",
	"tzfg: 11",
	"hbhh: cftr * fvcg",
	"fvcg: 3",
	"fhsf: wcgg + jtsj",
	"jlhc: nftd - fhns",
	"dsvh: vgjp - fwtf",
	"mhcw: 14",
	"dddr: fgjs + vmjs",
	"dnmm: 2",
	"gdqq: 1",
	"qzdn: 2",
	"hshq: 12",
	"htqd: vfps * ncrr",
	"hczz: 6",
	"cqwv: 16",
	"vpsr: rmtp * fgfs",
	"hwwc: 3",
	"wjrr: gntv * htdv",
	"cnvg: 4",
	"gfcs: 3",
	"zhss: gbbl + lcqg",
	"pmjj: 2",
	"wgmb: lpnt + cpqh",
	"pqlf: 3",
	"nbfp: 3",
	"pfqv: 5",
	"qhpn: jmfl - wmdq",
	"svcv: rgch + qljc",
	"vnqj: nlgj + mjmc",
	"jzfq: zdjs + chtz",
	"ctpp: lnvd * vgqg",
	"pnwl: 6",
	"smdn: hjqv * tdwc",
	"pdns: 10",
	"dqgc: brdd + brbr",
	"gmfs: nwtv * wpdl",
	"nfdb: mlwn * rbgr",
	"sjzb: 3",
	"wvwc: dhpv * clsb",
	"vlsh: 16",
	"jvbh: 2",
	"hfvm: 15",
	"npfq: 3",
	"vfjz: pvdl + rmnw",
	"plws: 11",
	"wspv: nqhd * jlrs",
	"bdjt: qzgv * czgv",
	"wlvh: 3",
	"tlvm: ctpp / tlfw",
	"szbh: 2",
	"zhpm: 4",
	"mbrd: sdwt * jvfl",
	"gcrs: fpsm * ntvl",
	"vmzn: 3",
	"frft: qdcf * qpqm",
	"qrfn: sdbg / fftc",
	"dfmj: jbsd * lblt",
	"ncgd: 2",
	"wrbj: cgvj + mhpf",
	"jggh: 2",
	"gmll: ggjg * mvbm",
	"vdns: tpjd + mpwq",
	"hwtt: 17",
	"zrgd: psgg * phqr",
	"vhrh: zzsp + rmrf",
	"cpjs: 2",
	"wcwt: 12",
	"wbzz: vjwd * jhnl",
	"dfdf: hgzt * hcfg",
	"lgfz: spjf + bqlg",
	"ntsl: cgnf + lwpz",
	"jjfw: 4",
	"dmfb: 2",
	"pvdl: 20",
	"jhzv: srdd + msmg",
	"cqjh: 17",
	"bqdv: 8",
	"mnzg: dmlq + wwbl",
	"mzgl: csnm + mlqd",
	"tnnv: 7",
	"lqmg: cffz + dhdq",
	"bcgj: bgfm + rgvg",
	"wbmt: 2",
	"qrvt: 3",
	"jtsj: sfhw * fmsb",
	"lbnn: wqdw + bqtz",
	"bptz: zvtz - bzln",
	"fcnr: pfpg + zmcs",
	"bflj: 7",
	"trhq: hnnj + grdt",
	"sbld: 3",
	"hpbb: htlt + npgq",
	"lcwp: 5",
	"dnrn: wpvg * hprw",
	"crjj: 2",
	"pfrq: 2",
	"ddqv: 3",
	"jzrp: 3",
	"ptvw: 1",
	"jjrm: 2",
	"zmdm: rtth * ssgz",
	"wqdw: 3",
	"vfps: lwnv + pmgc",
	"llcv: 6",
	"scpr: dgtd * njpp",
	"tzcq: 16",
	"gfhv: 2",
	"zldn: 3",
	"cpqh: 7",
	"vcdw: qwbn + gbcn",
	"hmvc: hhvw * bgqm",
	"ljls: 6",
	"grgc: 7",
	"zcpm: 14",
	"pqrl: wjll + rvbj",
	"mcgn: 3",
	"glgv: 1",
	"tgtt: plwv + grlf",
	"bgfm: 8",
	"clsb: mgcz * lmgh",
	"vwww: 3",
	"tltw: 3",
	"nzdm: 1",
	"cdgd: fhjr * lqjh",
	"hhpt: ctqd + wcdd",
	"lmfz: 2",
	"dvbf: dljc + cgtn",
	"vtwn: mvqt - vjdq",
	"lvlb: jwls + bntv",
	"cwhw: jfcs - vbcz",
	"rlql: 5",
	"vrtd: 3",
	"mpwq: hhzp + ndjl",
	"srpf: 3",
	"wlsj: 2",
	"lqhf: mqrj - zggt",
	"cbsg: zdqc * zfcg",
	"fgfs: 10",
	"npsq: 2",
	"gczr: 17",
	"fwvf: 6",
	"wdhv: hwjd * dhqb",
	"pgpt: 11",
	"cjlq: 2",
	"lwjb: 1",
	"tlsg: 18",
	"gthm: 2",
	"jdbl: ffvp * chlr",
	"nqdg: 4",
	"lnwn: 2",
	"pqlm: 5",
	"rjmr: ncwz + lvnp",
	"wnsb: lwjb + jlhc",
	"hrjr: zvbh + dtrm",
	"zszb: jmrv + pdbv",
	"jzlm: 2",
	"pzgp: zfrp * dhcd",
	"djqs: tddj * csvc",
	"bbvc: 5",
	"lhtg: 2",
	"hcfg: 6",
	"gdrh: pbmg * gdfl",
	"vhvb: rcdv / jzwh",
	"hhwr: cwnw + wjzc",
	"bwlw: 7",
	"hhvw: snnl * zldn",
	"hplj: pszm + tbgb",
	"cfmm: 6",
	"zfmm: vclq - vhvb",
	"ctsr: 6",
	"hcdp: 5",
	"fpzj: 5",
	"fqvr: tlvq + tmvg",
	"hqlm: zjpf * zsfw",
	"zwwn: 3",
	"wmrd: 3",
	"dfhj: brhs + snwv",
	"lsfr: 7",
	"zqqv: tqsv + fvhp",
	"wjqf: ndqv * tcrw",
	"vfqp: 3",
	"jgbj: 3",
	"pbcg: 2",
	"cfsq: hjlw + wlvh",
	"jwwl: 2",
	"ppqb: dcjd * vfqp",
	"fdmv: 19",
	"vmlz: 2",
	"wzhl: 10",
	"tqbm: vlvc * gsrj",
	"btcl: jgfp / bpmg",
	"cgcg: 19",
	"zszg: mlmw * bpvd",
	"prpm: 3",
	"pswb: rmbs * pnnw",
	"bswd: 3",
	"jfcs: tzhd + hqlm",
	"lbqv: whzc * cqjg",
	"sttl: 19",
	"mvlq: rdpq * jnsd",
	"bhgz: dpth * wdbd",
	"ssbw: tqbm / ztsh",
	"jlmz: mdbb + mgmt",
	"pgrf: nrww + zzgs",
	"lqnf: ghhb * qdlg",
	"hjpj: 6",
	"tjll: qmqn * djhz",
	"lwwp: wpdr * zlct",
	"grvv: qmsb + jbmr",
	"nthz: 5",
	"cfvj: vwlp * lwqv",
	"tsfl: 6",
	"frlv: hjpj * ldhv",
	"mndb: mgjp + qrcd",
	"jlhp: ztzd / hcqg",
	"sfvm: 4",
	"hhdm: 14",
	"qdvs: jrdq * pvdq",
	"bwdz: wzlz * vrgz",
	"jhqp: 7",
	"hzhg: plsl + vnqj",
	"nftd: ndnp * lgwl",
	"qsgv: 10",
	"bbhw: 3",
	"tjmf: 11",
	"qbsw: 3",
	"lbbz: zjlw + cftf",
	"tsnd: 5",
	"vfnl: mzgl + hsgj",
	"bjtf: 3",
	"jqmn: 2",
	"cjsv: 6",
	"ttws: vzbr + tlsg",
	"wsjm: 2",
	"cftf: 20",
	"qhbc: hmvq * vflz",
	"mvfq: cwth + vmbm",
	"pcvl: tctd * whqh",
	"glrh: ldct + hbhh",
	"slws: wglv * vvgm",
	"pgwq: dzgz + swls",
	"tcrw: 11",
	"jnrj: blqz * gqlw",
	"nlsh: ppqb + wbzz",
	"mrht: hzmm * fcpm",
	"pjfz: 9",
	"zhpb: qhpn / cdnb",
	"gcvq: 4",
	"mvqt: vtgv + vbcr",
	"mzlr: qwcr + nnmc",
	"hrdh: 2",
	"dpzw: 3",
	"qgsf: jggh * pgpt",
	"qhwd: 2",
	"nfbt: prpm * bzbq",
	"cvst: 2",
	"hmvn: 3",
	"lpnt: bhpf * vlpj",
	"mpjz: chzs * vfgb",
	"mzlw: 3",
	"mtmt: tlcq * thfw",
	"rnpw: zzdh + rfts",
	"pztv: jgbj * mdvz",
	"fdvs: 2",
	"bzbq: 2",
	"hprw: vbrw * dzzc",
	"hcss: 6",
	"jnsc: 5",
	"pszm: 8",
	"qtjw: svsv * lvrz",
	"scqb: nwlg + phsj",
	"hvvm: zdtb / lwgb",
	"htdv: dsbm / vsvt",
	"nhlf: 4",
	"mwmc: sjnv / jwwl",
	"cdnb: 2",
	"nnqw: 2",
	"cwth: sphd + mrht",
	"fszl: 8",
	"nlqq: mdph * pzjs",
	"vmts: thjd + pmcg",
	"vggg: vcgv * bwdz",
	"ttnp: 2",
	"zhff: bcwj * qtjh",
	"ntgn: qsgv + qwzs",
	"qsfh: 4",
	"bcwj: qgnm + mpph",
	"vwlp: 4",
	"ghhb: 2",
	"wqwj: 7",
	"scwv: 2",
	"jprd: 18",
	"lwnv: 2",
	"jzwh: 11",
	"thfw: ffvt + lbnn",
	"qjwn: rpws * nwgd",
	"phsj: 4",
	"lpqz: ptdh + mmbb",
	"gdlp: svrj * cnnd",
	"rnwb: zfgl + slrm",
	"csll: 4",
	"gwbf: 4",
	"whdl: 3",
	"rshl: 3",
	"swfj: dsvh * frss",
	"hvgv: tzcq * jhhv",
	"dsbm: pbtr * zqrr",
	"hnnj: cfdc * ztpp",
	"hsgj: hnff / dzcg",
	"bbfd: 14",
	"brdd: 6",
	"vgqh: ddqn * pfqv",
	"plsl: lwrc * vmzn",
	"vftl: 4",
	"jnmq: rclf + nszg",
	"wcdd: 9",
	"vlfh: mwmc + lbgh",
	"tlrh: bvlm + zmdm",
	"sfjp: hrjr + rgpp",
	"lmsn: 2",
	"wblz: bpbw + gjdl",
	"csqg: 9",
	"ndjl: 8",
	"cbcc: trgn * bpwg",
	"hqpm: pgrw + zgbf",
	"qszh: 2",
	"pvdq: gmfj * glrd",
	"jmnm: vhvm * zgsl",
	"wwbl: htrv * jnsc",
	"jchl: wtgr + gvqw",
	"cghq: 1",
	"svsv: 5",
	"lsqw: 2",
	"bfjd: hcsz * qsdn",
	"lmgh: 2",
	"tqmm: gbzv * pbpj",
	"nrww: dfdf + mpjm",
	"pdbv: pnwl * gthm",
	"vhvm: 3",
	"chtz: ffqn * bprt",
	"jbzr: 2",
	"dswn: 2",
	"cjtj: qdlr * vvwv",
	"tqrh: 4",
	"wdjr: hcdt * tfsh",
	"jrpj: 16",
	"fznq: 2",
	"jshq: pswb / qrwm",
	"vbrw: vtwv + fszl",
	"sfzg: htht + fcbq",
	"ccjp: 6",
	"dzdz: 3",
	"zscm: bqdw * mcgn",
	"mpch: vffd + mzjh",
	"cgvj: jdls / grqf",
	"crhv: rnpw + zqrj",
	"ccmc: 2",
	"fsht: zwgm * hwtt",
	"bqhh: gmbg * sbbn",
	"tlvq: hwff - jqhc",
	"gbpt: 2",
	"rwfv: 1",
	"qjtr: jdbl + rqjl",
	"bpqc: jvbh * wzhl",
	"vngp: 5",
	"thwh: 5",
	"dlct: 5",
	"pjgp: gmmz * zdwj",
	"gcqv: zcvw * lqgl",
	"bvlm: 2",
	"lhzq: mvbg * qwdn",
	"vqtw: 3",
	"lwrc: gtnz + mvfq",
	"fwgr: 3",
	"psgg: 17",
	"sbmp: lmjr + zcnb",
	"zdcj: 4",
	"zpmg: 2",
	"fpnj: 19",
	"vrqq: vjvd + qpdl",
	"hlbf: 2",
	"ddnn: nwgr + bfng",
	"pmsb: hgfd + hzhg",
	"glhq: 2",
	"tshz: tzdl + hvgf",
	"pbmg: shvd + npss",
	"snft: 2",
	"cnbj: 7",
	"snwv: 3",
	"djgq: 10",
	"sfwm: bzhs + fvtw",
	"vwws: gdtv * cmnw",
	"zczr: jdrf * czcd",
	"slcc: ssbw * wtbn",
	"ljdg: 11",
	"rcdv: mlzt - mtmt",
	"ccrm: 2",
	"wltd: 5",
	"jznt: 7",
	"htht: 19",
	"mqsv: sjsr * tvtm",
	"ffnd: pdsg / lntg",
	"jsmw: rshl * ddrp",
	"glrd: 4",
	"zjmr: 2",
	"pgrw: tzgm + vwcl",
	"qrtz: mfdp * spjh",
	"nflc: tsbz * zpmg",
	"sfrj: 15",
	"qqjz: 2",
	"vbcz: mwcg * ldqh",
	"gfhl: 3",
	"fsnc: 4",
	"hfbn: 1",
	"mgmt: 1",
	"pwtl: cvst * htwz",
	"svlz: jnth / nssq",
	"hlst: nzdm + ntnt",
	"zwzq: rcll + nmjd",
	"fhqd: 2",
	"dtpn: sntl * bdjq",
	"bnls: jbcp * vgfb",
	"jgtd: jvps + hcwr",
	"vzpd: mftt * dcls",
	"ftwv: 20",
	"rtbc: 2",
	"sfzf: vjdw - lqmg",
	"sjbm: 2",
	"zmpl: 2",
	"jgpg: qzqh * fpqf",
	"rhnr: jmwm * shvl",
	"jnhz: cqps * wftv",
	"wglv: 4",
	"wdzb: bmvz - rqjg",
	"dzjj: 2",
	"gmgc: 2",
	"zcqj: 4",
	"cmnc: 1",
	"tsfc: 2",
	"jzmv: djqs * clnd",
	"hprg: tsgs * gflr",
	"qpdl: ndgr * lbrr",
	"cpnl: 3",
	"gqrz: sjwm + lrcb",
	"bzcn: bfzm * nnzt",
	"gvqw: zhpj * twwc",
	"sdnq: 5",
	"jqhc: gdqq + jgtd",
	"jqtc: wspv + mlqs",
	"mlmc: trhq * hpjq",
	"ctcs: 18",
	"dqnv: 2",
	"lplj: 3",
	"qgbj: cqhf + ttrz",
	"rwmg: 3",
	"mlqd: mrns / nmjh",
	"wpdl: bdlp - lgfz",
	"rfqt: 5",
	"mlws: dwhb / dhmm",
	"vlpj: 4",
	"vzpl: 5",
	"lljf: 2",
	"fwhf: rwmg * dqgc",
	"nnzt: 7",
	"lgwl: 3",
	"tbwm: 1",
	"qwrh: crhv / jgpg",
	"fvnb: jhzv * jjnn",
	"nqzl: 17",
	"gtcd: vfnl + vvgc",
	"fftc: 6",
	"hhcs: dqwh * wqwj",
	"hwlw: grst * mpjp",
	"vztd: tgzt + scbg",
	"dmdt: 5",
	"dzms: 2",
	"vzzn: 18",
	"qwdn: srpf * vhwd",
	"vmgp: 7",
	"fcbr: stnv / rnsh",
	"jvbz: 2",
	"gpht: 5",
	"qmmd: 3",
	"gjpr: 6",
	"ftbj: 2",
	"mlqs: rfrj + rtfm",
	"jwhn: tzfg * gcrs",
	"cwdg: fdvs * rhff",
	"pqwp: wsfw + jscp",
	"mgcz: crgg + djgn",
	"ggrc: slvj * jvbz",
	"pfnj: 4",
	"zzgs: drtj * jqml",
	"ztzd: hnzt + bzcn",
	"nqwc: ftsh * fmfn",
	"wffs: gvfh + lsbg",
	"zlct: 2",
	"bwmp: 20",
	"lfdt: tvmd * zmpl",
	"nctl: 11",
	"tqvd: 5",
	"nqfw: 6",
	"bnqj: mzlr * lltp",
	"trjw: 2",
	"ffrf: 7",
	"tbzc: lwft - qrvt",
	"wbcs: wbpr + bwhw",
	"jnbr: dnrn + pvcc",
	"csnn: 2",
	"hlwc: 8",
	"vjsl: hczz + mzbw",
	"lwpz: 8",
	"jqgt: 3",
	"tsgs: 3",
	"sgjd: nctl * lqzr",
	"rqzc: htqd - bhvr",
	"hqdr: fhsf + cpjs",
	"wsns: 18",
	"rrcd: pwrg * hqqt",
	"gpfj: tgtt + lnhs",
	"vhwd: cbcc + zddf",
	"wzlz: 2",
	"mtbl: glrh * pplj",
	"fpls: 2",
	"nssq: 2",
	"hpdg: 3",
	"rbvs: 10",
	"gfqc: 2",
	"gvfh: dcwc + cmnc",
	"wllr: rsfg * llmr",
	"zscj: 14",
	"qtjh: 2",
	"nbbw: 19",
	"dcls: wltp * mzlw",
	"fvtw: hhqt / gtqd",
	"tfdm: 3",
	"zgtd: fnqr + lszb",
	"dnqd: dhcb + fwvf",
	"qgdc: 13",
	"twzl: 2",
	"zwds: nmsw + hbqq",
	"lcbc: pmsm + gvlt",
	"pvcc: slcc + frft",
	"wftv: 2",
	"qhpl: jzfq * pbdw",
	"wcgg: 5",
	"mrns: dvbf * qlzf",
	"npgq: 4",
	"cngp: rtmp + zdqg",
	"vgjp: rscb * pzlc",
	"qrcb: hdtw * pcjz",
	"tvgv: nszs * srwm",
	"hwjd: 5",
	"qmsb: 6",
	"lsfc: lwbh * nlpd",
	"lvgv: wwtm / jhbw",
	"wztz: 6",
	"pwlj: 1",
	"wztl: 3",
	"vbdh: 2",
	"vgqt: plcq * mnzg",
	"vcgv: 3",
	"njhs: 9",
	"lbvs: ttws + ffnd",
	"zrcq: 3",
	"vbrr: 5",
	"tzvp: 3",
	"jhnl: 5",
	"rpdq: 8",
	"lnhs: gqbt * jlmz",
	"nbcr: zczr + qdqq",
	"qrcd: dtpn / ttnp",
	"cftr: 7",
	"lzdq: pqrc * qltg",
	"qppn: 2",
	"whlb: 5",
	"mwjw: 4",
	"mdbb: dltt * cnvg",
	"lmnc: 11",
	"gflr: 5",
	"sbbn: mtdt + fbhc",
	"htmj: 5",
	"pnmf: gztd * pczj",
	"zrzd: 3",
	"pmgc: 5",
	"lrjw: 15",
	"cnnd: 7",
	"lwnf: zhff - zhss",
	"cpjr: 14",
	"wpdr: hcdp * qgvg",
	"mqsz: cwhw * qbzz",
	"chpg: 5",
	"ptdh: wfvb * tjll",
	"rqjl: 4",
	"blpd: 13",
	"jqsc: 20",
	"lrcd: 2",
	"mpjh: 9",
	"vgqg: 2",
	"bqdw: 3",
	"jlbb: tjfm + pqrl",
	"crhq: 2",
	"hlhl: cngp * dmfb",
	"mpph: hhdm * blff",
	"zwth: 2",
	"jmrl: nnfz * qhwd",
	"tsql: vbmz - hsvv",
	"htqf: ttsh * jthd",
	"bwvz: scms + bqrt",
	"mtst: hflz + rqzc",
	"lwgv: 7",
	"plwv: zglw * hhpc",
	"wdbd: 2",
	"zgdt: 11",
	"rddq: 4",
	"zdpc: 4",
	"nwtv: 3",
	"gdfl: 2",
	"tbgb: 1",
	"lmsm: bmbh + cnnj",
	"vvgc: vmvj * wsns",
	"hspv: wdhv + glvt",
	"mmbb: shzh / cnhr",
	"bwhw: nsdq + pmsb",
	"qhdr: 2",
	"hjwt: rnqb * dgsm",
	"bgqm: nwnt * pqwp",
	"zftn: 1",
	"cwrf: 5",
	"hpjq: 2",
	"msmg: zcpm + gfzj",
	"cffz: nbbw * pjql",
	"sbpt: 2",
	"clhp: hshq - mtlz",
	"vvtw: vjch * zvrb",
	"cvwh: 13",
	"hbsh: zmfs - rwsr",
	"dtzs: lgss + cqwv",
	"hsnp: 19",
	"lhhr: wfdb * pdns",
	"cpfg: 3",
	"fszt: njlj * ccmc",
	"vzhb: qrll + vzbq",
	"lmbj: hwwc * tnpp",
	"sptz: 12",
	"nqrv: pbgd + bqlv",
	"qrwm: 5",
	"blml: 7",
	"gtwz: 5",
	"vmvj: 5",
	"rdps: 3",
	"tmvv: 4",
	"llfz: 2",
	"rtdc: bhrr * lwfz",
	"qbnf: 13",
	"dvhn: 2",
	"nfbz: 5",
	"vsrr: 13",
	"mttf: jjvv - bjtf",
	"bgzf: zqqv * jpsf",
	"wwlp: bfjg + sblz",
	"bprt: tqdj * jgmn",
	"tlnd: 7",
	"pnnw: vlpp * hczs",
	"jttc: nhlf * bpfh",
	"jvrr: qpjd * vwww",
	"nnfz: 5",
	"lqgl: lwwp - lvlb",
	"zzsp: 6",
	"sphd: 16",
	"gprr: pmjj * cfsq",
	"fpmh: 9",
	"qljc: dtcr * nbwt",
	"fnqr: lfnh / ldmb",
	"ztmf: 4",
	"pfpg: qwpg + dwqw",
	"wbsm: tsql + qvvp",
	"dwqw: 2",
	"jmfl: fvzh * fhqd",
	"ntnt: 12",
	"drfl: 6",
	"vbgq: 2",
	"cmnw: 13",
	"dzhc: bpqc * dmdt",
	"cbvj: ddqv * ngsg",
	"gvlt: djwq * mrwf",
	"npzj: 5",
	"fdvh: 13",
	"jqml: hphh * ftbj",
	"zntp: mspm / wlsj",
	"snnl: 9",
	"qdtr: 10",
	"mhgh: 13",
	"smsg: fqjh + jpdg",
	"fdms: qmhd * fwbt",
	"ldcs: sgvh / sjbm",
	"mpjm: jgvf + jdbd",
	"lmrz: 12",
	"gsgv: gczr * wzvr",
	"mtlz: 2",
	"bgdg: 2",
	"ftjl: gpsn + mcwf",
	"znqt: 2",
	"tqnm: cbvj + sglt",
	"mclg: gfqc * rvvz",
	"lqgt: qmmd * qhdr",
	"sptt: pwtl + lhql",
	"wfdb: rnsn * fmnb",
	"wjcp: 3",
	"jvfl: 5",
	"lqqb: bfpz * lmsn",
	"bqsc: 3",
	"drjw: pjgp + rrdj",
	"gqbt: 3",
	"lrcb: 3",
	"fjjn: 7",
	"nczp: 2",
	"gzwd: 1",
	"mrjz: nqcc + clhp",
	"qpqm: wvqp + hhwr",
	"rlfq: cqfr * vsmm",
	"jwnl: 3",
	"cgtn: swjg * qdtr",
	"pwsf: ltpw + lqnf",
	"lwgb: 2",
	"jmwm: 2",
	"hcsz: 2",
	"npcz: smvl + gctm",
	"trtt: prwd * swqz",
	"mmwh: 2",
	"jwls: 4",
	"hnzd: htjq * nczp",
	"ztzl: czvz * fnld",
	"vlzm: ftwv * lwwf",
	"qbgf: lwmn + dcrn",
	"zhqv: plmq + mttn",
	"vmct: 2",
	"bsht: 2",
	"rcfn: pgrf + sfwm",
	"dqgj: 3",
	"lzmf: 3",
	"nggv: fqnl * llcv",
	"pdjl: 9",
	"fbzf: gdsl - cslv",
	"vbmz: fcnr / pbcg",
	"sjwm: gzwd + jvrr",
	"grlf: nfbz * zqdw",
	"cngs: msjj * cdgv",
	"fpfn: 11",
	"dhvm: 6",
	"glnr: 3",
	"vfls: 3",
	"rfwb: bqrw + dflq",
	"qlzf: 2",
	"pftc: 4",
	"sglt: lhtg * qgfj",
	"vshp: 4",
	"nplr: zmbh * bwlw",
	"nbwt: qzlv * crjj",
	"trsc: 8",
	"zmcs: 19",
	"gndq: vlsh * tdvn",
	"wrjw: 3",
	"vvgm: 2",
	"hbrc: 5",
	"lsrn: 2",
	"rgch: cctn * sdvg",
	"sjnv: bwvz + wvwc",
	"bfrr: svlz * vbmv",
	"tqzm: 3",
	"bvct: twzl * qhgg",
	"npdq: lqvh * vdns",
	"jnth: vfjz / bflj",
	"gqlw: 5",
	"gwmj: lsrn * jrpj",
	"mfcc: wncf + ztmf",
	"zmgr: qbsw + jjgj",
	"mqph: 3",
	"bvhg: zhqv * dgns",
	"pgtf: 4",
	"vjwd: 2",
	"nqhd: pdrz - jvfh",
	"wbqt: lhmp * flhd",
	"btws: wjrr / lbtr",
	"vjdq: bqdv * vzhb",
	"zrvj: 4",
	"sbpl: qrfn + mlmc",
	"zvtz: 16",
	"zmfs: tqrh * zcqj",
	"qzgv: qggr + dqcs",
	"dlgp: tscs - zctc",
	"pplj: 3",
	"jjbf: lfdt * vgnc",
	"ddmh: gjdc + cbvv",
	"lvrz: 5",
	"lcqg: zscm * rhrr",
	"qpjl: rnbz + dtmd",
	"wvqp: fdmv * ddlr",
	"rvvz: 5",
	"lqgr: tqnm * hpdg",
	"wmcb: 9",
	"vtwv: 5",
	"zgbf: 20",
	"gprz: 1",
	"tnpp: 3",
	"wvcz: vnbl + hlpj",
	"zcnb: tbdd * pfcm",
	"bmbh: ctsr * pqlm",
	"zfrp: 5",
	"ldhv: scpr + lzmj",
	"mcwf: mgcp * plsw",
	"frpn: 7",
	"plmq: dzms * bgpn",
	"zqrr: 2",
	"gldf: sppw + rmcd",
	"ngnr: 2",
	"zhpj: sqqn - rbpp",
	"nmjh: 2",
	"tgzt: 2",
	"zgzt: 13",
	"vnnj: 12",
	"jjnn: 10",
	"hzdn: hfnv + pjbj",
	"tlfw: 2",
	"qpfq: mwwz * pfnj",
	"msvr: 3",
	"lwwf: 2",
	"lqvh: 2",
	"qpch: twpr + wmcb",
	"rpdn: ndlw * llsw",
	"dzcg: 8",
	"tdlh: rspj * hmvn",
	"nszs: 3",
	"cdbj: rpnm * zthj",
	"vjrs: bsvd * mqlv",
	"svhn: cjtj + mhbb",
	"vjwj: rvdg - gprz",
	"hqqt: 2",
	"qhtc: 3",
	"plcq: 2",
	"hwfd: 5",
	"qghf: bgms * dbhd",
	"tdwc: 2",
	"glvt: hbqw + jrfd",
	"hzlr: 2",
	"wthc: fhlt + mclg",
	"lrfd: 2",
	"mjmc: lqhf + jprd",
	"dvjp: 8",
	"llrt: sqqg * vzvh",
	"ngft: jhqp + qdgd",
	"pcqr: cpfg + tvbs",
	"wpfs: fpfn * csll",
	"bbsz: 6",
	"lqzr: 5",
	"djwq: dczh + bflt",
	"nhtn: 1",
	"jdgl: 5",
	"jlrn: wztl * pcsr",
	"pmgn: 2",
	"bpqv: 2",
	"pcld: 13",
	"dhcb: 10",
	"mtzq: dlgp - wcwt",
	"pzjn: wcds * mzts",
	"bzhs: hvvm * ctgh",
	"htjq: lplj + wbsm",
	"dqtt: 4",
	"ndgr: 6",
	"rlln: 8",
	"hlzl: nnqw * zqls",
	"vcgj: 2",
	"hbqq: 5",
	"qgws: rtdc + dpcs",
	"hbqw: pzgp / qmbd",
	"rclf: 5",
	"lfnh: hspv * cggw",
	"dzgz: wpvd + tjmf",
	"bfpz: 3",
	"lzlv: 6",
	"lltp: 3",
	"hjgc: 18",
	"dlwp: 2",
	"rdls: cndh - jbqb",
	"pfcm: 4",
	"vfsl: 16",
	"srgg: hwlw * pjfz",
	"mvpp: gbgp + chgc",
	"hhqt: qwrh * njvl",
	"ldmb: 3",
	"snzv: 2",
	"pfhh: 2",
	"wmdq: ccjp * lrfd",
	"qmqn: 7",
	"lcph: 4",
	"tjzm: 4",
	"lblt: fcbr - wmlz",
	"wfvb: 3",
	"gsrj: 2",
	"whhd: 8",
	"jdrf: 3",
	"tsbz: 5",
	"zdjs: fpnj * bmzn",
	"vsmm: 3",
	"jpsf: tsfl + vnhj",
	"rwlb: hplj * cjlq",
	"wpvg: bjnn + btws",
	"vwcl: 18",
	"vnhj: 1",
	"bflt: fpls * hrdv",
	"jjrt: 2",
	"tswp: jdgl * vjrs",
	"whzc: 2",
	"smhj: 2",
	"hszd: ljls + bjng",
	"lqvl: qjwn * rhbl",
	"shzh: mdvr - sccb",
	"lljj: 3",
	"ddqn: 2",
	"tqsv: cqlj * shnq",
	"mlmw: 10",
	"vlpp: jhwm * dzdz",
	"dlsd: 5",
	"nmsw: 1",
	"ncwz: zwwn * jzlh",
	"mzts: 5",
	"gjdc: lbbz - zhpm",
	"rscb: vlfh * vngp",
	"nszg: 2",
	"wpwj: 3",
	"qtqw: jjrt * bjzq",
	"dtcr: jsbq + whlb",
	"mvsj: 2",
	"cwns: 4",
	"czvz: tmjw * cmth",
	"tmvg: 1",
	"sgvh: pcld * csdv",
	"ztbg: 4",
	"qmdh: tvss * vprg",
	"lbzp: gmch * bztw",
	"cnqn: 4",
	"wjzc: gqrz * sbld",
	"rqnh: 2",
	"splf: 8",
	"mvcb: qfvh / dvjp",
	"jbtc: npzj * wffw",
	"vvhc: sbmp + ptqv",
	"dwhb: rcrt * vbrr",
	"gnbw: wbcg + qjtr",
	"zgqh: hzlr * hrtf",
	"pmcg: prwl / vmct",
	"gbcn: ngjw + wvcz",
	"gqsn: 4",
	"pvpl: 11",
	"wcnc: 4",
	"nbjm: srhc * vbnc",
	"ddlr: 2",
	"qtlz: 8",
	"gfbn: 5",
	"jzjc: tdzq - nhtn",
	"qsgm: gjpr + ggbj",
	"mwqj: hlrd * vfls",
	"nphw: 2",
	"jrfd: vvtw + qgsf",
	"zrmn: 4",
	"qmhd: 11",
	"dcwc: 14",
	"shnq: 4",
	"wncf: 3",
	"pmsm: hbvm + mtnm",
	"fwnz: cshb + dnhn",
	"lzhh: 2",
	"hcsv: 8",
	"sjtl: htmj * pfhh",
	"mdph: qghf / vrtd",
	"dpcs: fpml + cvwh",
	"pzjj: 11",
	"bhrr: 4",
	"dcjd: hprg - nqdg",
	"hzmm: 2",
	"mvbm: 2",
	"ctgh: 2",
	"ncbb: 2",
	"zmvg: sdrn + hvmr",
	"mgws: 5",
	"tlbc: 3",
	"szgz: 2",
	"pwrg: szbh * ztzl",
	"glll: 3",
	"qfdv: 3",
	"zgsl: 3",
	"ggbj: ntgl - lljj",
	"rmtp: blwh - qcnw",
	"zcvw: 2",
	"dzrg: 3",
	"mdvz: 16",
	"zvbh: 3",
	"jjvv: zcqt * ccrm",
	"bpws: 4",
	"vfgb: 2",
	"rmrf: plws * msfm",
	"ngjw: mgws * gfbn",
	"mjdd: 3",
	"psnn: gmfs + qlgh",
	"gtnz: 5",
	"zcqt: zdcj + njhs",
	"clww: wdjr - csqg",
	"jmtf: 2",
	"nsns: hbsh + rffd",
	"zfgl: rwlb + wrbj",
	"bsmj: 3",
	"ssgz: 4",
	"pnld: gcqv + tdlh",
	"zdww: 2",
	"svsq: 4",
	"root: cgdh + qhpl",
	"mzdj: jzgf * dcll",
	"nmjd: 5",
	"sddj: 2",
	"wwgr: bwvw + nbjz",
	"cqps: cdbj - vmvg",
	"qzlv: 3",
	"mgjp: 12",
	"dwvt: bhgz * qjgd",
	"sdvg: hmhf - hbrc",
	"nqhn: 2",
	"zbgb: jjbf / jlsg",
	"zfvj: 15",
	"sfpn: sbpl + mtst",
	"cmhq: 2",
	"jgmn: 15",
	"jlsg: 2",
	"qnfh: wcft * bjfd",
	"hgdn: fcmg * qlgm",
	"ztsh: 2",
	"cndh: ncmh + jqsc",
	"vgwn: cbsg - hgdn",
	"rhrr: zrmn + fwgr",
	"vgrb: 2",
	"fqmt: 3",
	"wtdz: mfqf * bdvp",
	"twjf: mstb + hjgc",
	"mftt: hlbf + ctcs",
	"ftsh: 2",
	"hlpj: 7",
	"zpmp: 6",
	"fbwv: 3",
	"zglh: 3",
	"wgjb: 5",
	"vbcr: wrlb * zpdp",
	"bqtz: chpg * hcsv",
	"dtrm: 8",
	"tncp: 19",
	"mqlv: rdls * rqnh",
	"hcdl: 7",
	"sqwl: wjfb * smhj",
	"rpnm: 2",
	"jhwm: 3",
	"vzbr: ldzz * hdnl",
	"bjzq: tnrq * qppn",
	"dhzq: 4",
	"mzpv: wgjb * hsnp",
	"jdls: 16",
	"fzbj: zfmm / nfbt",
	"zpds: 2",
	"qdlr: 6",
	"mdnc: srpl * rjmr",
	"dmlq: tltw * splf",
	"gntv: 2",
	"bhpf: 4",
	"rvdg: tlvm + hqdr",
	"zbfl: vlbv + pwlj",
	"dtmd: sddj * flbh",
	"lhmp: 2",
	"hmvq: vgqt / lrcd",
	"lvnp: 5",
	"jrdq: 5",
	"gqmm: 2",
	"qdcd: 4",
	"cfdc: 2",
	"twps: qqjz * pcqr",
	"zprq: nlsh * vhps",
	"zdtb: zrgd * cgcg",
	"ndlw: zbzp * ntph",
	"cvmd: 7",
	"hnff: frlv + vvfg",
	"bhzp: 3",
	"tgrz: jmtf * lzdd",
	"fmfn: zflm + ncbb",
	"jwpp: fszt + bbvc",
	"lgrl: 3",
	"fwtf: wdvw * qdsb",
	"nwgd: lrrf + pgtf",
	"rnbz: fdgq * qnfh",
	"qsdn: tbzc * fscn",
	"hdtw: 4",
	"jhws: 2",
	"hwbg: mqhc + qdvs",
	"fmsb: 4",
	"hflz: prlz * vdqr",
	"nghf: jwnl * gnbw",
	"bjfd: 3",
	"vbmv: 3",
	"humn: 1104",
	"cqfr: 9",
	"svrm: 5",
	"fhlt: wcdz * npfq",
	"hdph: 2",
	"bztw: 4",
	"hrtf: mzpv + nlqq",
	"bhvr: cphz * dtsn",
	"mjrs: 3",
	"hlrd: dbch / crhq",
	"jgfp: glph + dwvt",
	"lggb: hlst * lnwn",
	"dhcd: hfbn + bqhh",
	"flhd: ntsl * wbmt",
	"jwpr: fqvr + lsvb",
	"njpp: 2",
	"qpjr: 8",
	"gjtr: 6",
	"rnfc: vphf + zftn",
	"prcm: wpfc + tqvd",
	"ctmf: 2",
	"sntl: lvgv * zdpc",
}
