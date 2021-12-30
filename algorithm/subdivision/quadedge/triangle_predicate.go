package quadedge

import "github.com/spatial-go/geoos/algorithm/matrix"

func isInCircleNormalized(a, b, c, p matrix.Matrix) bool {
	var (
		adx = a[0] - p[0]
		ady = a[1] - p[1]
		bdx = b[0] - p[0]
		bdy = b[1] - p[1]
		cdx = c[0] - p[0]
		cdy = c[1] - p[1]

		abdet = adx*bdy - bdx*ady
		bcdet = bdx*cdy - cdx*bdy
		cadet = cdx*ady - adx*cdy
		alift = adx*adx + ady*ady
		blift = bdx*bdx + bdy*bdy
		clift = cdx*cdx + cdy*cdy
	)

	disc := alift*bcdet + blift*cadet + clift*abdet
	return disc > 0
}

func isInCircleRobust(a,b,c,p matrix.Matrix) bool {
	return isInCircleNormalized(a, b, c, p)
}
