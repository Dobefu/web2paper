package converter

type Rect struct {
	Height float32
	Width  float32
}

type PdfSize Rect

var (
	PdfSizeA0        = Rect{Width: 2383.94, Height: 3370.39}
	PdfSizeA1        = Rect{Width: 1683.78, Height: 2393.94}
	PdfSizeA2        = Rect{Width: 1190.55, Height: 1683.78}
	PdfSizeA3        = Rect{Width: 841.89, Height: 1190.55}
	PdfSizeA4        = Rect{Width: 595.28, Height: 841.89}
	PdfSizeA5        = Rect{Width: 419.53, Height: 595.28}
	PdfSizeA6        = Rect{Width: 297.64, Height: 419.53}
	PdfSizeA7        = Rect{Width: 209.76, Height: 297.64}
	PdfSizeA8        = Rect{Width: 147.40, Height: 209.76}
	PdfSizeA9        = Rect{Width: 104.88, Height: 147.40}
	PdfSizeA10       = Rect{Width: 73.70, Height: 104.88}
	PdfSizeB0        = Rect{Width: 2834.65, Height: 4008.19}
	PdfSizeB1        = Rect{Width: 2004.09, Height: 2834.65}
	PdfSizeB2        = Rect{Width: 1417.32, Height: 2004.09}
	PdfSizeB3        = Rect{Width: 1000.63, Height: 1417.32}
	PdfSizeB4        = Rect{Width: 708.66, Height: 1000.63}
	PdfSizeB5        = Rect{Width: 498.90, Height: 708.66}
	PdfSizeB6        = Rect{Width: 354.33, Height: 498.90}
	PdfSizeB7        = Rect{Width: 249.45, Height: 354.33}
	PdfSizeB8        = Rect{Width: 175.75, Height: 249.45}
	PdfSizeB9        = Rect{Width: 124.72, Height: 175.75}
	PdfSizeB10       = Rect{Width: 87.87, Height: 124.72}
	PdfSizeC0        = Rect{Width: 2599.37, Height: 3676.54}
	PdfSizeC1        = Rect{Width: 1836.85, Height: 2599.37}
	PdfSizeC2        = Rect{Width: 1298.27, Height: 1836.85}
	PdfSizeC3        = Rect{Width: 918.43, Height: 1298.27}
	PdfSizeC4        = Rect{Width: 649.13, Height: 918.43}
	PdfSizeC5        = Rect{Width: 459.21, Height: 649.13}
	PdfSizeC6        = Rect{Width: 323.15, Height: 459.21}
	PdfSizeC7        = Rect{Width: 229.61, Height: 323.15}
	PdfSizeC8        = Rect{Width: 161.57, Height: 229.61}
	PdfSizeC9        = Rect{Width: 113.39, Height: 161.57}
	PdfSizeC10       = Rect{Width: 79.37, Height: 113.39}
	PdfSizeRA0       = Rect{Width: 2437.80, Height: 3458.27}
	PdfSizeRA1       = Rect{Width: 1729.13, Height: 2437.80}
	PdfSizeRA2       = Rect{Width: 1218.90, Height: 1729.13}
	PdfSizeRA3       = Rect{Width: 864.57, Height: 1218.90}
	PdfSizeRA4       = Rect{Width: 609.45, Height: 864.57}
	PdfSizeSRA0      = Rect{Width: 2551.18, Height: 3628.35}
	PdfSizeSRA1      = Rect{Width: 1814.17, Height: 2551.18}
	PdfSizeSRA2      = Rect{Width: 1275.59, Height: 1814.17}
	PdfSizeSRA3      = Rect{Width: 907.09, Height: 1275.59}
	PdfSizeSRA4      = Rect{Width: 637.80, Height: 907.09}
	PdfSizeExecutive = Rect{Width: 521.86, Height: 756.00}
	PdfSizeLegal     = Rect{Width: 612.00, Height: 1008.00}
	PdfSizeLetter    = Rect{Width: 612.00, Height: 792.00}
	PdfSizeTabloid   = Rect{Width: 792.00, Height: 1224.00}
)
