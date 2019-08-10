# hbar.awk -- generate horizontal bar chart from name\tvalue pairs
BEGIN {
	if (top == 0) {
		top=90
	}
	if (right == 0) {
		right=80
	}
	if (left == 0) {
		left=20
	}
	if (linesize == 0) {
		linesize = 2.5
	}
	if (step == 0) {
		step = linesize*1.2
	}
	if (color == "") {
		color = "lightsteelblue"
	}
	if (labelsize == 0) {
		labelsize = 1.5
	}
	if (labelcolor == "") {
		labelcolor = "gray"
	}
	lx=left-2
	minlen=0.5
	maxlen=right-left
	print "deck"
}
NR == 1 { 
	min = max = $2
}
{
	label[NR] = $1
	data[NR] = $2
	if ($2 > max) {
		max = $2
	}
}
END {
	min=0
	printf "\tlx=%g\n",				lx
	printf "\tlabelsize=%g\n",		labelsize
	printf "\tlabelcolor=\"%s\"\n", labelcolor
	printf "\tcolor=\"%s\"\n",		color
	printf "\tleft=%g\n",			left
	printf "\tminlen=%g\n",			minlen
	printf "\tmaxlen=%g\n",			maxlen
	printf "\tlinesize=%g\n",		linesize
	print  "\tslide"
	print "\t\tclist 50 95 3 \"sans\" \"\" 100 1.2"
	print "\t\t\tli \"ACRES OF LAND OWNED BY NEGROES\""
	print "\t\t\tli \"IN GEORGIA.\""
	print "\t\telist"
	printf "\t\tctext \"338,769\" 22 %g 2\n", top-(linesize/4)
	printf "\t\tctext \"1,062,223\" 50 5.3 2\n"

	for (i=1; i <=NR; i++) {
		printf "\t\tv%d=vmap %g %g %g minlen maxlen\n", i, data[i], min, max
		vn++
	}
	y=top
	for (i=1; i<=NR; i++) {
		printf "\t\tetext \"%s\" lx %g labelsize \"serif\" labelcolor\n", label[i], y-(labelsize/4)
		printf "\t\thline left %g v%d linesize color\n", y, i
		y-=step
	}
	print "\teslide"
	print "edeck"
}
