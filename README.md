# BranchComparer

# For using as plugin:

# Build plugin
# go build -buildmode=plugin -o branchcomparer.so branchcomparer.go
# 2. Copy to the your library with shared libs
# sudo cp ./lib.so /usr/local/lib
# 3. Update cashe of shared libs
# sudo ldconfig

# Go-code for using library:
    //Создаем интерфейс для использования динамической библиотеки
    type Comparer interface {
        Compare()
    }

    // загружаем плагин
    // 1. открываем .so файл для загрузки символов
    plug, err := plugin.Open({Place/Name.so})
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // 2. выполняем поиск символов(экспортированных функций или переменных)
    // в нашем случае, это ExtComparer
    symcomparer, err := plug.Lookup("ExtComparer")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // 3. делаем возможным работу с этими символами в нашем коде
    // нужно не забывать про тип экспортированных символов
    var comparer Comparer
    comparer, ok := symcomparer.(Comparer)
    if !ok {
        fmt.Println("unexpected type from module symbol")
        os.Exit(1)
    }

    // 4. используем загруженный плагин
    comparer.Compare()