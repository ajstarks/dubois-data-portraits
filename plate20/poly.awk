BEGIN {
	aspect=612.0/792.0
}
NR == 1 {
	minx=$1
	maxx=$1

	miny=$2
	maxy=$2
}
NF >= 2 {
	x[NR]=$1
	y[NR]=$2

	if ($1 < minx) {
		minx = $1
	}
	if ($1 > maxx) {
		maxx = $1
	}
	if ($2 < miny) {
		miny = $2
	}
	if ($2 > maxy) {
		maxy = $2
	}
}

# Map one range to another
function vmap(value, low1, high1, low2, high2) {
	return low2 + (high2-low2)*(value-low1)/(high1-low1)
}
function minvalue(a) {
	minimum=a[1]
	for (i=1; i < length(a); i++) {
		if (a[i] < minimum) {
			minimum = a[i]
		}
	}
	return minimum
}

function maxvalue(a) {
	maximum=a[1]
	for (i=1; i < length(a); i++) {
		if (a[i] > maximum) {
			maximum = a[i]
		}
	}
	return maximum
}

function poly(left, right, bottom, top) {
	if (vminx != 0) {
		minx = vminx
	}
	if (vmaxx != 0) {
		maxx = vmaxx
	}
	if (vminy != 0) {
		miny = vminy
	}
	if (vmaxy != 0) {
		maxy = vmaxy
	}
	print "#", minx, maxx, miny, maxy
	printf "polygon \""
	pminx=10000
	pmaxx=-10000
	for (i=1; i < NR; i++) {
		px=vmap(x[i], maxx, minx, left, right)
		if (px > pmaxx) {
			pmaxx=px
		}
		if (px < pminx) {
			pminx=px
		}
		printf "%g ", px
	}
	printf "%g\"", vmap(x[1], maxx, minx, left, right)
	printf " \t\t\""

	pminy=10000
	pmaxy=-10000
	for (i=1; i < NR; i++) {
		py=vmap(y[i], miny, maxy, bottom, top)
		if (py > pmaxy) {
			pmaxy=py
		}
		if (py < pminy) {
			pminy=py
		}
		printf "%g ", py
	}
	printf "%g\" \"%s\" 80\n", vmap(y[1], miny, maxy, bottom, top), color
	printf "ctext \"%s\" %g %g 1\n", landvalue, pminx+((pmaxx-pminx)/2), pminy+((pmaxy-pminy)/2)
}

END {
	poly(5, 80, 15, 80)
}
