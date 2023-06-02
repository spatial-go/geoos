package buffer

import (
	"testing"

	"github.com/spatial-go/geoos"
	"github.com/spatial-go/geoos/algorithm/matrix"
)

func TestBuffer(t *testing.T) {
	type args struct {
		geom     matrix.Steric
		distance float64
		quadsegs int
	}
	tests := []struct {
		name string
		args args
		want matrix.Steric
	}{
		{name: "point buffer", args: args{
			geom:     matrix.Matrix{100, 90},
			distance: 50,
			quadsegs: 4,
		}, want: matrix.PolygonMatrix{
			{{150, 90}, {146.193976625564, 70.8658283817455}, {135.355339059327, 54.6446609406727},
				{119.134171618255, 43.8060233744357}, {100, 40}, {80.8658283817456, 43.8060233744356}, {64.6446609406727, 54.6446609406725},
				{53.8060233744357, 70.8658283817454}, {50, 89.9999999999998}, {53.8060233744356, 109.134171618254}, {64.6446609406725, 125.355339059327},
				{80.8658283817453, 136.193976625564}, {99.9999999999998, 140}, {119.134171618254, 136.193976625564}, {135.355339059327, 125.355339059328},
				{146.193976625564, 109.134171618255}, {150, 90}},
		},
		},

		{name: "line buffer", args: args{
			geom:     matrix.LineMatrix{{100, 100}, {300, 300}},
			distance: 50,
			quadsegs: 4,
		}, want: matrix.PolygonMatrix{

			{{264.6446609406726, 335.3553390593274}, {280.8658283817456, 346.19397662556435}, {300.00000000000006, 350}, {319.1341716182545, 346.19397662556435}, {
				335.3553390593274, 335.3553390593274}, {346.19397662556435, 319.1341716182545}, {350, 300.00000000000006}, {346.19397662556435, 280.8658283817456}, {
				335.3553390593274, 264.6446609406726}, {135.35533905932738, 64.64466094067262}, {119.13417161825444, 53.80602337443564}, {99.99999999999997, 50}, {
				80.8658283817455, 53.80602337443566}, {64.64466094067262, 64.64466094067262}, {53.806023374435675, 80.86582838174549}, {50, 99.99999999999994}, {
				53.80602337443563, 119.13417161825441}, {64.64466094067262, 135.35533905932738}, {264.6446609406726, 335.3553390593274}},
		},
		},

		{name: "line buffer1", args: args{
			geom:     matrix.LineMatrix{{12925610.710013045, 4856270.465739663}, {12925562.342807498, 4856370.568506251}, {12925554.8522816, 4856387.370610655}},
			distance: 10,
			quadsegs: 8,
		}, want: matrix.PolygonMatrix{

			{{12925553.33876128, 4856366.217971605}, {1.2925553209322736e+07, 4.856366496718426e+06}, {12925545.718796838, 4856383.29882283}, {12925545.099927789, 4856385.158915575}, {
				12925544.855836228, 4856387.104002521}, {12925544.995902454, 4856389.059335069}, {12925545.514743807, 4856390.949770884}, {12925546.392421499, 4856392.702661579}, {
				12925547.59520687, 4856394.250644547}, {12925549.076877553, 4856395.534231671}, {12925550.780493775, 4856396.504095417}, {12925552.64058652, 4856397.122964467}, {
				12925554.585673466, 4856397.367056029}, {12925556.541006014, 4856397.2269898}, {12925558.431441829, 4856396.70814845},
				{12925560.184332523, 4856395.830470758}, {12925561.732315492, 4856394.627685386}, {12925563.015902616, 4856393.1460147025}, {12925563.985766362, 4856391.44239848},
				{12925571.47629226, 4856374.640294076},
				{12925619.714059263, 4856274.816274309}, {12925620.389796244, 4856272.976077729}, {12925620.693544587, 4856271.039410265}, {12925620.61363141, 4856269.080696963}, {12925620.153127734, 4856267.175210074},
				{12925619.329730455, 4856265.396176391}, {12925618.175082272, 4856263.811963183}, {12925616.733555663, 4856262.483450873}, {12925615.06054769, 4856261.461693445},
				{12925613.220351111, 4856260.785956463}, {12925611.283683648, 4856260.482208122}, {12925609.324970344, 4856260.562121298}, {12925607.419483457, 4856261.022624974}, {12925605.640449774, 4856261.846022252},
				{12925604.056236565, 4856263.000670436}, {12925602.727724256, 4856264.442197044}, {12925601.705966827, 4856266.115205017}, {12925553.33876128, 4856366.217971605}},
		},
		},

		{name: "poly buffer", args: args{
			geom:     matrix.PolygonMatrix{{{100, 100}, {100, 200}, {200, 200}, {200, 100}, {100, 100}}},
			distance: 50,
			quadsegs: 4,
		}, want: matrix.PolygonMatrix{
			{{100, 50}, {80.86582838174527, 53.80602337443576}, {64.64466094067251, 64.64466094067274}, {53.80602337443563, 80.86582838174559}, {50, 100}, {50, 200}, {
				53.80602337443566, 219.1341716182545}, {64.64466094067262, 235.35533905932738}, {80.86582838174552, 246.19397662556435}, {100, 250}, {200, 250}, {
				219.1341716182545, 246.19397662556435}, {235.35533905932738, 235.35533905932738}, {246.19397662556435, 219.1341716182545}, {250, 200}, {250, 100}, {
				246.19397662556435, 80.86582838174552}, {235.35533905932738, 64.64466094067262}, {219.1341716182545, 53.80602337443566}, {200, 50}, {100, 50}},
		},
		},

		{name: "multi point buffer", args: args{
			geom:     matrix.Collection{matrix.Matrix{100, 100}, matrix.Matrix{200, 200}},
			distance: 50,
			quadsegs: 4,
		}, want: matrix.PolygonMatrix{
			{{150, 100}, {146.19397662556435, 80.86582838174553}, {135.3553390593274, 64.64466094067265}, {119.13417161825454, 53.80602337443568}, {100.00000000000009, 50}, {
				80.86582838174562, 53.806023374435625}, {64.64466094067271, 64.64466094067254}, {53.80602337443571, 80.86582838174539}, {50, 99.99999999999984}, {
				53.80602337443559, 119.13417161825431}, {64.64466094067248, 135.35533905932724}, {80.86582838174532, 146.19397662556426}, {99.99999999999977, 150}, {
				119.13417161825426, 146.19397662556443}, {135.35533905932718, 135.35533905932758}, {146.19397662556423, 119.13417161825477}, {150, 100}},
		},
		},

		{name: "polygon buffer", args: args{
			geom: matrix.PolygonMatrix{{{122.993197, 41.117725}, {122.999399, 41.115696}, {122.99573, 41.109516},
				{122.987146, 41.106994}, {122.984775, 41.107699}, {122.990687, 41.117878}, {122.993197, 41.117725}}},
			distance: 0.001,
			quadsegs: 4,
		}, want: matrix.PolygonMatrix{{{122.99325784324384, 41.118723147333654}, {122.99350793587271, 41.11867543089336}, {122.99970993587272, 41.11664643089336},
			{123.00002331519273, 41.11647717254184}, {123.0002574970166, 41.11620881855708}, {123.00038277421193, 41.11587541098052}, {123.00038325474164, 41.11551924422623},
			{123.00025887764812, 41.11518549982346}, {122.99658987764812, 41.10900549982346}, {122.99634342934228, 41.10872625039283}, {122.99601188795101, 41.1085565526679},
			{122.98742788795101, 41.1060345526679}, {122.98686098957647, 41.106035475582736}, {122.98448998957647, 41.10674047558273}, {122.98416932802418, 41.10690328556774},
			{122.98392699858715, 41.10716900603418}, {122.98379434195499, 41.10750327110905}, {122.98378851473228, 41.10786284998192}, {122.98391027055912, 41.108201237985504},
			{122.98982227055912, 41.118380237985505}, {122.99006123336798, 41.11865801033471}, {122.99038421204416, 41.11883105794881}, {122.99074784324384, 41.11887614733365},
			{122.99325784324384, 41.118723147333654}},
		},
		},
		{name: "poly buffer", args: args{
			geom:     matrix.PolygonMatrix{{{100, 100}, {200, 100}, {200, 200}, {100, 200}, {100, 100}}},
			distance: 50,
			quadsegs: 4,
		}, want: matrix.PolygonMatrix{
			{{50, 100}, {53.80602337443563, 80.86582838174559}, {64.64466094067251, 64.64466094067274}, {80.86582838174527, 53.80602337443576}, {100, 50}, {200, 50},
				{219.1341716182545, 53.80602337443566}, {235.35533905932738, 64.64466094067262}, {246.19397662556435, 80.86582838174552}, {250, 100}, {250, 200},
				{246.19397662556435, 219.1341716182545}, {235.35533905932738, 235.35533905932738}, {219.1341716182545, 246.19397662556435}, {200, 250}, {100, 250},
				{80.86582838174552, 246.19397662556435}, {64.64466094067262, 235.35533905932738}, {53.80602337443566, 219.1341716182545}, {50, 200}, {50, 100}},
		},
		},
		{name: "issue87", args: args{
			geom: matrix.PolygonMatrix{{{12695002.300434208, 2578061.0407920154}, {12695011.873910416, 2577943.999133326},
				{12695182.081411814, 2577948.097385522}, {12695210.133923491, 2578044.286065328}, {12695102.042697946, 2578089.728780503},
				{12695002.300434208, 2578061.0407920154}}},
			distance: 108,
			quadsegs: 8,
		}, want: matrix.PolygonMatrix{
			{{1.2694972447614357e+07, 2.578164832935971e+06}, {1.2694952801158171e+07, 2.5781570294458857e+06}, {1.2694935051145962e+07, 2.5781455483847535e+06},
				{1.2694919877626132e+07, 2.5781308296214156e+06}, {1.2694907861935072e+07, 2.578113437069374e+06}, {1.2694899464424662e+07, 2.5780940370818204e+06}, {1.2694895006825024e+07, 2.578073372921938e+06},
				{1.269489465991821e+07, 2.57805223628659e+06}, {1.2694904233394418e+07, 2.5779351946279006e+06}, {1.2694907808887677e+07, 2.577915111920888e+06}, {1.2694915096776368e+07, 2.5778960597306876e+06},
				{1.269492583707376e+07, 2.5778787177214357e+06}, {1.2694939646632573e+07, 2.577863704548606e+06}, {1.269495603281328e+07, 2.577851555789213e+06}, {1.2694974411058454e+07, 2.5778427048357427e+06},
				{1.269499412574614e+07, 2.5778374674353916e+06}, {1.2695014473578395e+07, 2.577836030426163e+06},
				{1.2695184681079794e+07, 2.577840128678359e+06}, {1.2695207182815429e+07, 2.5778430549201253e+06}, {1.2695228576484643e+07, 2.5778506181148915e+06},
				{1.2695247917693771e+07, 2.5778624843959813e+06}, {1.2695264352652138e+07, 2.5778781299429825e+06}, {1.2695277155861448e+07, 2.577896864105062e+06}, {1.26952857621419e+07, 2.5779178598887883e+06},
				{1.2695313814653577e+07, 2.5780140485685943e+06}, {1.2695317824740773e+07, 2.5780361198009993e+06}, {1.2695317188721681e+07, 2.578058543350649e+06}, {1.2695311934036078e+07, 2.5780803517980585e+06},
				{1.2695302287386933e+07, 2.578100604261106e+06}, {1.2695288664959753e+07, 2.578118426987522e+06}, {1.2695271654467084e+07, 2.578133051051196e+06}, {1.2695251989792826e+07, 2.5781438455259646e+06},
				{1.269514389856728e+07, 2.5781892882411396e+06}, {1.26951205234274e+07, 2.578196135841541e+06}, {1.2695096208269283e+07, 2.578197571070177e+06}, {1.2695072189878095e+07, 2.5781935209244587e+06}, {1.2694972447614357e+07, 2.578164832935971e+06}},
		},
		},
	}
	for _, tt := range tests {
		if !geoos.GeoosTestTag &&
			tt.name != "line buffer1" {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := Buffer(tt.args.geom, tt.args.distance, tt.args.quadsegs); got == nil || !got.EqualsExact(tt.want, 0.5) {
				t.Errorf("Buffer() = %v,\n want %v", got, tt.want)
			}
		})
	}
}
