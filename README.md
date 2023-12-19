# BranchComparer

For using as plugin:
* Build plugin
  >go build -buildmode=plugin -o branchcomparer.so branchcomparer.go
* Copy to the your library with shared libs
  >sudo cp ./lib.so /usr/local/lib
* Update cashe of shared libs
  >sudo ldconfig

Go-code for using library:

    // Создаем интерфейс для использования динамической библиотеки
    type Comparer interface {
        Compare(b1 string, b2 string) (string, error)
    }

    // Загружаем плагин
    // Открываем .so файл для загрузки символов
    plug, err := plugin.Open({Place/Name.so})
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Выполняем поиск символов(экспортированных функций или переменных)
    symcomparer, err := plug.Lookup("ExtComparer")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // Делаем возможным работу с этими символами в нашем коде
    var comparer Comparer
    comparer, ok := symcomparer.(Comparer)
    if !ok {
        fmt.Println("unexpected type from module symbol")
        os.Exit(1)
    }

    // Используем загруженный плагин
    comparer.Compare(branch1, branch2)
