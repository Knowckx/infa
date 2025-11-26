# In-fa
In-fa is just a collection of utils of golang

Repo was created after Typhoon In-fa tracking inland south of Shanghai in July 2021.



# 解决拉不到新版本的问题
go get -u github.com/Knowckx/infa@3d2ed45baa3c93c4b6c15c0b642ac56b0ce43952


直接引用本地的版本 这样不会有缓存
go mod edit -replace github.com/Knowckx/infa=../infa