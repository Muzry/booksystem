function createORUpdateBookInfo(id){
    var hostURL
    var methodType
    if (id){
        hostURL = "http://127.0.0.1:8080/api/v1alpha1/books/" + id
        methodType = "PUT"
    } else {
        hostURL = "http://127.0.0.1:8080/api/v1alpha1/books"
        methodType = "POST"
    }
    $.ajax({
        type: methodType,
        dataType: "json",
        contentType: "application/json",
        url: hostURL,
        data: JSON.stringify(getFormData('#bookInfo')),
        success: function (data) {
            window.location = "./book.html"
        },
        error : function() {
            alert("异常！");
        }
    });
}


function createORUpdatePublisherInfo(id){
    var hostURL
    var methodType
    if (id){
        hostURL = "http://127.0.0.1:8080/api/v1alpha1/publishers/" + id
        methodType = "PUT"
    } else {
        hostURL = "http://127.0.0.1:8080/api/v1alpha1/publishers"
        methodType = "POST"
    }
    $.ajax({
        type: methodType,
        dataType: "json",
        contentType: "application/json",
        url: hostURL,
        data: JSON.stringify(getFormData('#publisherInfo')),
        success: function (data) {
            window.location = "./publisher.html"
        },
        error : function() {
            alert("异常！");
        }
    });
}

function cancleButton(html){
    window.location = "./" + html
}

function deletePublisherInfo(id){
    if (confirm("确定删除？")){
        $.ajax({
            type: 'DELETE',
            url: "http://127.0.0.1:8080/api/v1alpha1/publishers/" + id,
            dataType : "json",
            async: true,
            crossDomain: true,
            success:function(data) {
                getPublisherList()
            },
            error:function() {
                //error
            }
        })    
    }
}

function deleteBookInfo(id){
    if (confirm("确定删除？")){
        $.ajax({
            type: 'DELETE',
            url: "http://127.0.0.1:8080/api/v1alpha1/books/" + id,
            dataType : "json",
            async: true,
            crossDomain: true,
            success:function(data) {
                getBookList()
            },
            error:function() {
                //error
            }
        })    
    }
}

function getBookList(){
    $.ajax({
        type: 'GET',
        url: "http://127.0.0.1:8080/api/v1alpha1/books",
        dataType : "json",
        async: true,
        crossDomain: true,
        success:function(data) {
            for (var i = data.length - 1; i >= 0; i--) {
                data[i]['createAt'] = moment(data[i]['createAt']).format('YYYY年MM月DD日 HH:mm:ss')
            }
            newData = {
                list: data
            }
            var html = template('bookTable', newData)
            document.getElementById('table').innerHTML = html;
            feather.replace();
        },
        error:function() {
            //error
        }
    })
}

function getPublisherList(){
    $.ajax({
        type: 'GET',
        url: "http://127.0.0.1:8080/api/v1alpha1/publishers",
        dataType : "json",
        async: true,
        crossDomain: true,
        success:function(data) {
            for (var i = data.length - 1; i >= 0; i--) {
                data[i]['createAt'] = moment(data[i]['createAt']).format('YYYY年MM月DD日 HH:mm:ss')
            }
            newData = {
                list: data
            }
            var html = template('publisherTable', newData)
            document.getElementById('table').innerHTML = html;
            feather.replace();
            
        },
        error:function() {
            //error
        }
    })
}

function getBookrInfo(){
    var request = window.location.search.substr(1)
    var info = request.split("=")
    if (info.length == 2) {
        if (info[0] == 'id') {
            $.ajax({
                type: 'GET',
                url: "http://127.0.0.1:8080/api/v1alpha1/books/" + info[1],
                dataType : "json",
                async: true,
                crossDomain: true,
                success:function(infoData) {
                    $.ajax({
                        type: 'GET',
                        url: "http://127.0.0.1:8080/api/v1alpha1/publishers",
                        dataType : "json",
                        async: true,
                        crossDomain: true,
                        success:function(publisherData) {
                            newData = {
                                list: publisherData,
                                info: infoData,
                                title: "更新图书信息",
                            }
                            var html = template("bookInfoScript",newData)
                            document.getElementById('bookInfoDiv').innerHTML = html;
                            // document.getElementById('publisherOption').innerHTML = html;
                            feather.replace();
                        },
                        error:function() {
                            //error
                        }
                    })
                },
                error:function() {
                    //error
                }
            })
        }
    } else{
        $.ajax({
        type: 'GET',
        url: "http://127.0.0.1:8080/api/v1alpha1/publishers",
        dataType : "json",
        async: true,
        crossDomain: true,
        success:function(data) {
            newData = {
                list: data,
                title: "创建图书信息",
                info: {    
                    "id": "",
                    "barCode": "",
                    "name": "",
                    "author": "",
                    "publisher": "",
                    "pageNumber": "",
                    "size": "",
                    "version": "",
                    "edition": "",
                    "price": "",
                    "location": "",
                    "annotation": ""
                }
            }
            var html = template("bookInfoScript",newData)
            document.getElementById('bookInfoDiv').innerHTML = html;
            // document.getElementById('publisherOption').innerHTML = html;
            feather.replace();
        },
        error:function() {
            //error
        }
    })
    }  
}


function getPublisherInfo(){
    var request = window.location.search.substr(1)
    var info = request.split("=")
    if (info.length == 2) {
        if (info[0] == 'id') {
            $.ajax({
                type: 'GET',
                url: "http://127.0.0.1:8080/api/v1alpha1/publishers/" + info[1],
                dataType : "json",
                async: true,
                crossDomain: true,
                success:function(data) {
                    
                    newData = {
                        info: data,
                        title: "更新出版社信息",
                    }
                    var html = template("publisherInfoScript",newData)
                    document.getElementById('publisherInfoDiv').innerHTML = html;
                    feather.replace();
                },
                error:function() {
                    //error
                }
            })
        }
    } else{
        newData = {
            title: "创建出版社信息",
            info: {"isbn":"", "name": "", "id": ""},
        }
        var html = template("publisherInfoScript",newData)
        document.getElementById('publisherInfoDiv').innerHTML = html;
        feather.replace();
    }  
}

function getFormData(form) {
    var unindexed_array = $(form).serializeArray();
    var indexed_array = {};

    $.map(unindexed_array, function (n, i) {
      indexed_array[n['name']] = n['value'];
    });

    return indexed_array;
}