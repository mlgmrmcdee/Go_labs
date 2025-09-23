module laba1

go 1.25.1

//Powershell команда для сборки каждого .go файла в .exe
// Get-ChildItem -Recurse -Filter *.go |
// Where-Object { (Get-Content $_.FullName -Raw) -match 'package\s+main' } |
// ForEach-Object {
//     $dir  = $_.DirectoryName
//     $name = [IO.Path]::GetFileNameWithoutExtension($_.Name)
//     go build -o (Join-Path $dir "$name.exe") $_.FullName
// }