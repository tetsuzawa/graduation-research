package adf

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/gonum/floats"
	"github.com/tetsuzawa/go-research/adflib/misc"
)

func TestFiltRLS_Run(t *testing.T) {
	rand.Seed(1)
	//creation of data
	//number of samples
	n := 64
	L := 4
	//input value
	var x = make([][]float64, n)
	//noise
	var v = make([]float64, n)
	//desired value
	var d = make([]float64, n)
	var xRow = make([]float64, L)
	for i := 0; i < n; i++ {
		xRow = misc.Unset(xRow, 0)
		xRow = append(xRow, rand.NormFloat64())
		x[i] = append([]float64{}, xRow...)
		v[i] = rand.NormFloat64() * 0.1
		d[i] = x[i][0]
	}

	type args struct {
		d []float64
		x [][]float64
	}
	tests := []struct {
		name    string
		args    args
		want    []float64
		want1   []float64
		want2   [][]float64
		wantErr bool
	}{
		{
			name: "Run RLS Filter",
			args: args{
				d: d,
				x: x,
			},
			want:    []float64{-0.3916108209122005, 0.7295028042517872, -0.6069536686049593, 0.4367452376665853, 0.23711830894235797, 0.6307099188599508, -1.0230725554658542, 0.8328042966922731, -0.0345823704649979, 0.12182107719048674, -0.1438856650081231, 0.8802507001774842, -0.3817966179152634, -0.08279243237359318, 0.931368921270108, -1.1124290688507559, 0.36943731338011165, -0.36644659071077407, 0.14557273499634477, -0.31148652899834944, 0.7021839269774173, -0.8002252269760273, 0.5364409768096253, -0.4729094496945779, 0.16618349413014571, 0.4293084370375395, 0.3043448592028126, -0.17370505826416982, 0.6648047372541294, -1.0223548618367746, 0.24489390769174865, 0.3071038546891367, -0.7464448209883984, 0.38053853435028706, 0.5345519216822019, -0.5154949563504854, 0.3615051439491534, 0.3388634380826506, -0.6519205198043806, 0.31383848656752533, -1.1647942358154428, 0.4005490880425621, 0.1777475494646767, 0.2668615029400295, 0.4611044787158347, -1.1051470569314823, 0.24276928680981186, -0.24789889500801957, -0.19740882989083613, -0.28760618339371824, 0.9120479077956632, 0.5604369484957844, -0.2553247306536982, -0.44799569857959803, 0.5602303053124302, 0.8396035826915077, 0.0829316436906902, 0.25468695643856926, -0.6458871039878162, 0.5597885157656031, 0.9956036244697702, -0.010238859892847813, 0.5144997883929885, -0.07031503869557743,},
			want1:   []float64{0.3916108209122005, -0.7295028042517872, 0.6069536686049593, -1.6705034152645322, -0.7581128800955083, -0.3079046662483709, 1.1818802956422898, -1.5640873128697521, 1.619986332745621, 1.1770197703269476, 0.8763275908126362, -0.18012979773749938, 1.381422738926526, -0.23374481052049506, 0.1693602724799128, 2.1021394890592875, -1.8044842355123398, 0.5037969642841249, -0.991667108451942, 0.4676130242811641, -0.4224402083142458, 1.5029754885382085, -1.6059863791626459, 0.8031988353072397, -1.2868975864867789, 0.4979900337265437, 0.6686895226403726, 0.683821168750186, -0.41326469319498016, 1.1890455807550722, -1.883534985548183, 0.8471269590209005, -0.024116126861395415, -1.1567517442240542, 0.8857275751066782, 0.1850381938755158, -0.6085398607793895, 1.396769818571911, 0.4460865925475207, -1.3655037841714694, 0.48727068955230357, -2.3621869991274704, 1.8081717325164555, -0.307999611009022, 0.7426271357705825, 0.2665542944121322, -1.645862015183477, 1.2221820588966412, -1.7955535401006464, 0.6688399236936687, -0.5024767132362136, 1.9698014948566223, -1.014917987841795, 0.02736788595827455, -0.5580696367097926, 1.543256762740433, 0.10833447410991745, -0.24468075845373033, 0.187574255781577, -1.036246094731792, 1.7355182465546761, 0.0627094985401174, 0.2853766137605771, 0.7956525266392888,},
			want2:   [][]float64{{0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,}, {0.3911225845957826, 0.22868724321870512, -0.13139422199649511, 0.10084567175345303,},},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			af := Must(NewFiltRLS(L, 1.0, 1e-5, "random"))
			got, got1, got2, err := af.Run(tt.args.d, tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() got = %v\n, want %v\n", got, tt.want)
				for i := 0; i < n; i++ {
					fmt.Printf("%g, ", got[i])
				}
				fmt.Println("")
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Run() got1 = %v\n, want %v\n", got1, tt.want1)
				for i := 0; i < n; i++ {
					fmt.Printf("%g, ", got1[i])
				}
				fmt.Println("")
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Run() got2 = %v\n, want %v\n", got2, tt.want2)
				for i := 0; i < n; i++ {
					fmt.Print("{")
					for k := 0; k < L; k++ {
						fmt.Printf("%g, ", got2[i][k])
					}
					fmt.Print("}, ")
				}
				fmt.Println("")
			}
		})
	}
}

func ExampleExploreLearning_rls() {
	rand.Seed(1)
	//creation of data
	//number of samples
	n := 64
	L := 4
	mu := 1.0
	eps := 0.001
	//input value
	var x = make([][]float64, n)
	//noise
	var v = make([]float64, n)
	//desired value
	var d = make([]float64, n)
	var xRow = make([]float64, L)
	for i := 0; i < n; i++ {
		xRow = misc.Unset(xRow, 0)
		xRow = append(xRow, rand.NormFloat64())
		x[i] = append([]float64{}, xRow...)
		v[i] = rand.NormFloat64() * 0.1
		d[i] = x[i][L-1]
	}

	af, err := NewFiltRLS(L, mu, eps, "random")
	check(err)
	es, mus, err := ExploreLearning(af, d, x, 0.001, 1.0, 100, 0.5, 100, "MSE", nil)
	check(err)

	res := make(map[float64]float64, len(es))
	for i := 0; i < len(es); i++ {
		res[es[i]] = mus[i]
	}
	//for i := 0; i < len(es); i++ {
	//	fmt.Println(es[i], mus[i])
	//}
	eMin := floats.Min(es)
	fmt.Printf("the step size mu with the smallest error is %.3f\n", res[eMin])
	//output:
	//the step size mu with the smallest error is 0.798
}