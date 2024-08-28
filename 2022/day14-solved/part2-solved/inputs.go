package day

// ------------------------------ inputs ------------------------------------//
var testInput = []string{
	"498,4 -> 498,6 -> 496,6",
	"503,4 -> 502,4 -> 502,9 -> 494,9",
}
var expected = 93

var realInput = []string{
	"513,151 -> 513,155 -> 510,155 -> 510,161 -> 519,161 -> 519,155 -> 515,155 -> 515,151",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"488,36 -> 488,32 -> 488,36 -> 490,36 -> 490,26 -> 490,36 -> 492,36 -> 492,26 -> 492,36 -> 494,36 -> 494,32 -> 494,36",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"510,167 -> 514,167",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
	"486,120 -> 486,122 -> 482,122 -> 482,128 -> 491,128 -> 491,122 -> 490,122 -> 490,120",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
	"488,36 -> 488,32 -> 488,36 -> 490,36 -> 490,26 -> 490,36 -> 492,36 -> 492,26 -> 492,36 -> 494,36 -> 494,32 -> 494,36",
	"498,175 -> 498,176 -> 505,176",
	"493,39 -> 493,43 -> 489,43 -> 489,51 -> 501,51 -> 501,43 -> 496,43 -> 496,39",
	"489,79 -> 493,79",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"483,85 -> 487,85",
	"486,120 -> 486,122 -> 482,122 -> 482,128 -> 491,128 -> 491,122 -> 490,122 -> 490,120",
	"480,82 -> 484,82",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"494,144 -> 494,145 -> 504,145 -> 504,144",
	"486,54 -> 490,54",
	"488,36 -> 488,32 -> 488,36 -> 490,36 -> 490,26 -> 490,36 -> 492,36 -> 492,26 -> 492,36 -> 494,36 -> 494,32 -> 494,36",
	"492,60 -> 496,60",
	"493,39 -> 493,43 -> 489,43 -> 489,51 -> 501,51 -> 501,43 -> 496,43 -> 496,39",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"513,151 -> 513,155 -> 510,155 -> 510,161 -> 519,161 -> 519,155 -> 515,155 -> 515,151",
	"493,39 -> 493,43 -> 489,43 -> 489,51 -> 501,51 -> 501,43 -> 496,43 -> 496,39",
	"497,147 -> 497,148 -> 513,148 -> 513,147",
	"493,39 -> 493,43 -> 489,43 -> 489,51 -> 501,51 -> 501,43 -> 496,43 -> 496,39",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"488,36 -> 488,32 -> 488,36 -> 490,36 -> 490,26 -> 490,36 -> 492,36 -> 492,26 -> 492,36 -> 494,36 -> 494,32 -> 494,36",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
	"513,151 -> 513,155 -> 510,155 -> 510,161 -> 519,161 -> 519,155 -> 515,155 -> 515,151",
	"474,70 -> 478,70",
	"466,103 -> 466,104 -> 481,104 -> 481,103",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"471,85 -> 475,85",
	"490,131 -> 490,134 -> 482,134 -> 482,139 -> 501,139 -> 501,134 -> 495,134 -> 495,131",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"488,36 -> 488,32 -> 488,36 -> 490,36 -> 490,26 -> 490,36 -> 492,36 -> 492,26 -> 492,36 -> 494,36 -> 494,32 -> 494,36",
	"493,39 -> 493,43 -> 489,43 -> 489,51 -> 501,51 -> 501,43 -> 496,43 -> 496,39",
	"480,70 -> 484,70",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
	"468,70 -> 472,70",
	"480,60 -> 484,60",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"477,68 -> 481,68",
	"483,79 -> 487,79",
	"495,85 -> 499,85",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"493,39 -> 493,43 -> 489,43 -> 489,51 -> 501,51 -> 501,43 -> 496,43 -> 496,39",
	"486,120 -> 486,122 -> 482,122 -> 482,128 -> 491,128 -> 491,122 -> 490,122 -> 490,120",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"474,82 -> 478,82",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
	"490,131 -> 490,134 -> 482,134 -> 482,139 -> 501,139 -> 501,134 -> 495,134 -> 495,131",
	"495,63 -> 499,63",
	"489,63 -> 493,63",
	"477,79 -> 481,79",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"501,170 -> 505,170",
	"497,147 -> 497,148 -> 513,148 -> 513,147",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
	"489,85 -> 493,85",
	"488,36 -> 488,32 -> 488,36 -> 490,36 -> 490,26 -> 490,36 -> 492,36 -> 492,26 -> 492,36 -> 494,36 -> 494,32 -> 494,36",
	"490,131 -> 490,134 -> 482,134 -> 482,139 -> 501,139 -> 501,134 -> 495,134 -> 495,131",
	"490,131 -> 490,134 -> 482,134 -> 482,139 -> 501,139 -> 501,134 -> 495,134 -> 495,131",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"471,68 -> 475,68",
	"488,36 -> 488,32 -> 488,36 -> 490,36 -> 490,26 -> 490,36 -> 492,36 -> 492,26 -> 492,36 -> 494,36 -> 494,32 -> 494,36",
	"486,76 -> 490,76",
	"483,73 -> 487,73",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"513,151 -> 513,155 -> 510,155 -> 510,161 -> 519,161 -> 519,155 -> 515,155 -> 515,151",
	"466,103 -> 466,104 -> 481,104 -> 481,103",
	"504,167 -> 508,167",
	"507,170 -> 511,170",
	"488,36 -> 488,32 -> 488,36 -> 490,36 -> 490,26 -> 490,36 -> 492,36 -> 492,26 -> 492,36 -> 494,36 -> 494,32 -> 494,36",
	"486,120 -> 486,122 -> 482,122 -> 482,128 -> 491,128 -> 491,122 -> 490,122 -> 490,120",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"513,151 -> 513,155 -> 510,155 -> 510,161 -> 519,161 -> 519,155 -> 515,155 -> 515,151",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"507,164 -> 511,164",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"486,120 -> 486,122 -> 482,122 -> 482,128 -> 491,128 -> 491,122 -> 490,122 -> 490,120",
	"474,66 -> 478,66",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
	"486,120 -> 486,122 -> 482,122 -> 482,128 -> 491,128 -> 491,122 -> 490,122 -> 490,120",
	"488,36 -> 488,32 -> 488,36 -> 490,36 -> 490,26 -> 490,36 -> 492,36 -> 492,26 -> 492,36 -> 494,36 -> 494,32 -> 494,36",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
	"486,120 -> 486,122 -> 482,122 -> 482,128 -> 491,128 -> 491,122 -> 490,122 -> 490,120",
	"489,57 -> 493,57",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
	"488,36 -> 488,32 -> 488,36 -> 490,36 -> 490,26 -> 490,36 -> 492,36 -> 492,26 -> 492,36 -> 494,36 -> 494,32 -> 494,36",
	"480,76 -> 484,76",
	"488,36 -> 488,32 -> 488,36 -> 490,36 -> 490,26 -> 490,36 -> 492,36 -> 492,26 -> 492,36 -> 494,36 -> 494,32 -> 494,36",
	"477,63 -> 481,63",
	"492,82 -> 496,82",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
	"513,170 -> 517,170",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
	"490,131 -> 490,134 -> 482,134 -> 482,139 -> 501,139 -> 501,134 -> 495,134 -> 495,131",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"483,57 -> 487,57",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"498,175 -> 498,176 -> 505,176",
	"486,60 -> 490,60",
	"497,147 -> 497,148 -> 513,148 -> 513,147",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"513,151 -> 513,155 -> 510,155 -> 510,161 -> 519,161 -> 519,155 -> 515,155 -> 515,151",
	"486,82 -> 490,82",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"490,131 -> 490,134 -> 482,134 -> 482,139 -> 501,139 -> 501,134 -> 495,134 -> 495,131",
	"477,85 -> 481,85",
	"493,23 -> 493,21 -> 493,23 -> 495,23 -> 495,18 -> 495,23 -> 497,23 -> 497,19 -> 497,23 -> 499,23 -> 499,18 -> 499,23 -> 501,23 -> 501,18 -> 501,23 -> 503,23 -> 503,15 -> 503,23 -> 505,23 -> 505,16 -> 505,23",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"483,63 -> 487,63",
	"513,151 -> 513,155 -> 510,155 -> 510,161 -> 519,161 -> 519,155 -> 515,155 -> 515,151",
	"466,103 -> 466,104 -> 481,104 -> 481,103",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"463,98 -> 463,97 -> 463,98 -> 465,98 -> 465,94 -> 465,98 -> 467,98 -> 467,94 -> 467,98 -> 469,98 -> 469,88 -> 469,98 -> 471,98 -> 471,90 -> 471,98 -> 473,98 -> 473,97 -> 473,98 -> 475,98 -> 475,95 -> 475,98",
	"494,144 -> 494,145 -> 504,145 -> 504,144",
	"493,39 -> 493,43 -> 489,43 -> 489,51 -> 501,51 -> 501,43 -> 496,43 -> 496,39",
	"490,131 -> 490,134 -> 482,134 -> 482,139 -> 501,139 -> 501,134 -> 495,134 -> 495,131",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
	"494,144 -> 494,145 -> 504,145 -> 504,144",
	"476,117 -> 476,114 -> 476,117 -> 478,117 -> 478,110 -> 478,117 -> 480,117 -> 480,110 -> 480,117 -> 482,117 -> 482,114 -> 482,117 -> 484,117 -> 484,113 -> 484,117 -> 486,117 -> 486,113 -> 486,117",
}
