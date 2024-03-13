package day16

var testInput = []string{
	`.|...\....`,
	`|.-.\.....`,
	`.....|-...`,
	`........|.`,
	`..........`,
	`.........\`,
	`..../.\\..`,
	`.-.-/..|..`,
	`.|....-|.\`,
	`..//.|....`,
}

var expected = 46
var expected2 = 51

var realInput = []string{
	`\...|....||/............\.......-.......................\............./..........--......................./...`,
	`.......-..........................................-..-......../....../........................\...............`,
	`............/....\................-................-....................|..................\..../.............`,
	`.............|.\.................\.................|..................|..-........-...................-\......`,
	`..|...................|.....|......................./|...../......-................/.-..|.........\...........`,
	`.|......-..........\....../....\............\........../..\....|......................................-.......`,
	`....\....|/......../............\..............................-....../\|..-..\../.........\..\|..............`,
	`.........-..........-.........../...................-...........|....|\.......................................`,
	`.....|........................-......./...\../......\./.|................\......|.-...-........|.|/\..........`,
	`././............|............|...........\.....\..../...|....|...-.....|..............|................-......`,
	`................/../....\..........\......-....................-../............../........../-....\...........`,
	`......|........-................/........-......./........|......-\.............../........../.....|..........`,
	`-\.\......|...........//...\-.../..\..-..\.............................../............../.\-.-.......\........`,
	`.\..-...|./......./..............................\.................-....\../.................\..../...........`,
	`-....../......|....-...||...|.................../...-........-................|.......-/......\\...|..........`,
	`...../.........\..............\.................\..|....../....\....../.-\.....................|..\...........`,
	`\|..................\................-........./.......|...........\...../.................../................`,
	`..................../......\..........\..-...............\........-......................................|-...`,
	`................-................../.................................................-..................\-....`,
	`..\...................|....../........................./....-..........|.....................-................`,
	`..............-...................................\...............|./...-../.-|........||.....................`,
	`...........................././...|..........\............|...-....-..../.\\.......\...-.........\......../...`,
	`....\........................\..........-............\/...........|..-..\.....................................`,
	`..../.........../-\..\.........................|.........-........................|...........................`,
	`...|./..\...............................\....................../..............................................`,
	`..-.................\....|......-...-.........\-|..........\.-......|.....\.\....|......................|.....`,
	`......-.........\.....|..............-......................\.\...|\.....-..-..../............................`,
	`.....-..././......\|............................\.-.....................-.....................................`,
	`/-....................................-....................../.................................../.........\\.`,
	`.........|./...................................-..........|..\...\........../\.....\.....|...\.....|../..|....`,
	`............|.....................-.....................--.|-............................\|......../.|.....-..`,
	`-......\.........................\........|.\......../....||..||....................-........|.-.-............`,
	`................././/|............../..\.........-|.....-...|....|.......................-.......|............`,
	`|......|....../..............-............/........................-........\....\....|...../..............-..`,
	`.............//..../.......-........\.............|...../............................/........\......|........`,
	`...................-......-..........\..............\..............................\...................\......`,
	`...............\./......|.-...................-.|.........\.............../................/..................`,
	`.\..........-........|....................-..................||...........|...................................`,
	`....................\......../.........|............-...........-.........\..............-.-......../...\.....`,
	`........-.........|............\/.................-./.............-.......\...................-...............`,
	`....................-...............-./.................|............./.......\.............\........./\......`,
	`......\...............|........-/..................\.............../|.......\.|..............--..|.....|......`,
	`......\...\.....././.........\||.............................|....\.\.............../....-..|/........-.......`,
	`....../...../.................\...........|./.......|...\./....-...\........../.\......./..........-........\.`,
	`..................\..../....|................................................................|.....|....-.....`,
	`|..............................\......-................/|./-......./...........\........-.................|...`,
	`.\..|\....-..............-..|...........\........|....|..|........./.................................../......`,
	`...../...|..||.....\.--..........-....../............\..\.....|./..\....................\|......./............`,
	`.......\...|......-.......|..../......\..\/.|./.\..........|....|..............-.......................|...../`,
	`..............-...........\..\../...................\.........|.|....../..-/..............-...................`,
	`.......-.......\..../................................................|...../....|................|...-........`,
	`......................./........................|..\...|...-..-.........../....../.......-....|........../....`,
	`.......|..................................-.....\-./.\.........................\..../....-...|................`,
	`........../.../........./...........-.......|......./.-.........................\...\-.....................-..`,
	`..........\.\.............|........\..|........./.................../...........\.................\.../.......`,
	`....................-...|......|................../--.-.............-......................./.................`,
	`......./.........|......./../.....................-....//..|..............-..|.............\............\.....`,
	`...|...........-........................../.-..\............................\...|................\.//.......\.`,
	`../.....\..\.......-.............../.....|../.......................\-.................../.......-|/..........`,
	`......\.....................|.....\.............-.....\........../..........\.....-............|....\...-.....`,
	`...../.....................-|...\..-......./.............../.|..../................\..-...........\........./.`,
	`......|..\...........................\.........../.-............|....................-....|-.......-....\.....`,
	`........................................./.......|......-../../.............../....................|..........`,
	`........\.....|....../.......\.........../....../....-.........-.............||\...................-...|......`,
	`...........|.................\..\.............|...............\.....|..............................|..........`,
	`............|....../\....\..../.................\.-./......../\|../|..........................................`,
	`........................................................-|....................\....|.............|..-.........`,
	`/.|.............\\......\......./\.....................|.../.....................-...-.........../.-.....\....`,
	`.........................\/....-...\......./......-...........................--..|........\-..../............`,
	`.............-|...\........./..\........../..../..-.......\\.............................../.......|..../|....`,
	`....-...../......|.................|........-.......|.....|.../........-.....\.../..\.................../.....`,
	`.............../................|...........\............/-..|......../............-./.-./...............-....`,
	`.....\............|..\...............\..............\\.....-......./|....|.....--../...-..............\.......`,
	`................|............../................/......./............-..-.................-....||......./....\`,
	`.\..............\................|.......|.............|.........\|........\....-..../.............\..\/......`,
	`./.....\.......|......|.\..|.../.............-......\...............................-......\......\...........`,
	`.....|...../........../.......|.......................-.-.....\..................\............/....--./..-....`,
	`....................|..............-.\\......\..........\.....\.-.../.-.......-.......-.............\......-..`,
	`........./...................|.../....\...|...............|./../...-.........\-.................-.............`,
	`....................../.....\....-.\.............|............-......\.........\......\.......................`,
	`................\..\...|....-./..............-...\\.........\..........|......-.|/./.|..........\.\......../..`,
	`......\........-.............-/.............../...............................-.............|..........|/.....`,
	`....|.......-|...................................../.........\................./..........................\...`,
	`....\.......\\........-....\........\....\.............\............\........./........................\......`,
	`../.............................................................-......\..............|...................../.`,
	`............/../.........................\................................................../...-.............`,
	`.\........................\/.-...................-.......-...............|.....................-..............`,
	`./........................|.|/................\..............././..................................|.-........`,
	`............/........-..-..........|..............-........\../.|...\............................../..........`,
	`............\/......-.........\\............-.....................\..........\....................|..../..../.`,
	`..........\.../.....\.....................-..\..\.|..|....../-..|......-...|............-.....................`,
	`......................-........\.........-......................|..........-..................................`,
	`.|................\......-...........|.../....../............................................/......../...|../`,
	`.........-..........\...-...-................\............/..-....|\..............-.......\.|..-.../-.........`,
	`..|....../......./../-.............../....-......|-................................./......-...-...........-..`,
	`.............../............../....-....\......../.\../...................-............../\............|.\.|..`,
	`..|....\................-.....|.......................|\......................\.\..-...-..-...........|.......`,
	`.......................|......................\.../......|.................-..../..-..-.\....\................`,
	`..........................................-..........|-......................./-............-.................`,
	`..../...../....\.|........................................././.../...........................\................`,
	`......./...................................-............|-./.-.........\.|.../.-//...................-........`,
	`/......................|............|/.............../-..............-...............|.......|......|...|.....`,
	`....|.....-....-.....................\../..-...................\........./.../................./......-.......`,
	`........................\..............................-.\.-...|..................../............./.....-..../`,
	`......-.-...-........-.......\........../......|.........|..................../.....\./.....--...............|`,
	`.........-/...............-\...\..........\................|........./.../........\.|......|..................`,
	`..|/.................../......................................................../.....\..|.....|......-....-..`,
	`......................./......................\.........-./......................-............................`,
	`.....-.|.\................-............................|.............-........................../....-........`,
	`..../..|..|.\....|....\.................\.........|...-..-................/......|............/..\..........|.`,
}