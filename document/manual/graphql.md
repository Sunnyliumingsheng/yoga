# graphql

## 我想使用到graphql来写一些模块

graphql是一个用来代替restapi的,比较适合用来减轻前后端对接压力,另外对需求的修改提供了很好的支持   
目前来看restapi在应对多个前端的时候还是比较舒服的,因为不同的前端使用到的语言不一样,可能是C#,golang,js,vue等等   
但是开发者管理工具只用到了golang语言,可以尝试着使用一下   

网站位置
https://psj.codes/go-graphql-go

## 正文

graphQL 是fb开发的,并捐赠给了linux基金会   
并且相比于rest优势是巨大的
根本原因是其精准   
-----

一般来说有两个操作:
- query
- mutation

举例来说
一个book对象,查询条件为id

    query {
    book(id: 123) {
        title
        genre
        author {
        name
        }
    }
    }

一个createBook操作,指定了title,genre,author,并且要求返回id

    mutation {
    createBook(title: "The Ink Black Heart", genre: "Mystery", author: "J. K. Rowling") {
        id
        title
        genre
        author {
        id
        name
        }
    }
    }


-----

到这里我有点想放弃了

graph似乎有点太难了,首先是golang的包大,学习成本高,而且我个人开发的话不存在那种前后端解耦的需求,这个技术也不太稳定   
graphql本质上还是让前后端解耦,前端甚至可以不用和后段交流就能去写代码,甚至能前后端同步开发
可是我是一个人,没有这个需求啊    
而且gl有些后段被牵着鼻子走的感觉，我必须让已有的数据库去迁就这个


我的想法是从查询类的模块，使用graphql，如果涉及到用户信息处理这种强逻辑的功能,尽量还是使用rest   
涉及到前端调用后端主要是crud的业务模块可以使用gl   

比如查看现在有哪些课程