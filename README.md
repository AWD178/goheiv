Simple wrapper for https://github.com/monostream/tifig in golang

<h3>Usage</h3>
<code>
HEIV.SetFile("path_to_heic").SetOutput("output_path_(jpg, gif, png)").Convert()
</code>

Convert return cmd output or errors

<code>
result, error := HEIV.SetFile("path_to_heic").SetOutput("output_path_(jpg, gif, png)").Convert()
fmt.Println(result, err)
</code>

<h3>Resize image</h3>
<code>
HEIV.SetFile("path_to_heic").SetOutput("output_path_(jpg, gif, png)").SetWidth(YOUR_WIDTH).SetHeight(YOUR_HEIGHT).Convert()
</code>

<h3>Crop images</h3>

<code>
HEIV.SetFile("path_to_heic").SetOutput("output_path_(jpg, gif, png)").SetWidth(YOUR_WIDTH).SetHeight(YOUR_HEIGHT).UseCrop().Convert()
</code>
