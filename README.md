# CrappyPortScanner

	Really ropey port scanner in Go ( POC!! )
	1 - This just fires up loads of noisy threads, an IDS will see this coming a mile off!
	2 - The routine handling is crap, just waiting off the end for it to be done. needs some WGs
	3 - Only does tcp right now and only does up to ports 1024...very slowly!
	4 - My first attempt at something useful using channels and routines!
