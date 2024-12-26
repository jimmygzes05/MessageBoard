$(document).ready(function () {
    $.ajax({
        type: "get",
        url: "../session",
        success: function (response) {
            sessionStorage.setItem("userName", response.userName)
            sessionStorage.setItem("userID", response.userID)
            sessionStorage.setItem("nickName", response.nickName)
            $('#hiUser').text(`你好！${sessionStorage.getItem("nickName")}`)
        }
    });
})

$(document).ready(function () {
    $.ajax({
        type: "get",
        url: "../api/comment",
        success: function (response) {
            const fragment = $(document.createDocumentFragment());
            $.each(response.data, function (index, comment) {
                const card = $(`
                    <div class="card bg-default">
                        <span class="card-header d-flex justify-content-between">
                            <span class="text-primary">${comment.nickname}</span>
                            <span class="text-secondary">${comment.CreatedAt}</span>
                        </span>
                        <div class="card-body">
                            <p class="card-text" id="p_comment_${comment.id}">${comment.content}</p>
                            <form id="updateForm_${comment.id}" class="d-none">
                                <div class="form-group pb-1">
                                    <textarea class="form-control" name="content" rows="3">${comment.content}</textarea>
                                </div>
                                <button type="button" onclick="updateComment(${comment.id})" class="btn btn-sm btn-outline-success">發佈</button>
                            </form>
                        </div>
                        ${comment.userID == parseInt(sessionStorage.getItem('userID'), 10) ? `
                            <div class="card-footer pl-1">
                                <button type="button" class="btn btn-link btn-sm text-info pl-1 d-inline" onclick="showUpdate(${comment.id})">
                                    修改
                                </button>
                                <button type="button" class="btn btn-link btn-sm text-danger pl-1 d-inline" onclick="deleteComment(${comment.id})">
                                    刪除
                                </button>
                            </div>
                        ` : ''}
                    </div>
                `)[0];
                
                // 將卡片元素加入 DocumentFragment
                fragment.append(card);
            });
            $('#commentForm').after(fragment).after('<br>');
        }
    });
})

$('#logout').click(function (e) {
    $(window).attr('location', 'login');
    e.preventDefault();
});

$('#postButton').click(function (e) {
    var originData = $("#commentForm").serializeArray();
    userData = { name: "userID", value: parseInt(sessionStorage.getItem('userID'), 10) };
    nameData = { name: "nickName", value: sessionStorage.getItem('nickName') };
    originData[1] = userData;
    originData[2] = nameData;
    data = JSON.stringify(convertJSON(originData))
    console.log(data);
    $.ajax({
        type: "post",
        url: "../api/comment",
        data: data,
        headers: {
            'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
        },
        success: function (response) {
            // console.log(response);
            location.reload();
        }
    });
    e.preventDefault();
});

showUpdate = function showUpdate(id) {
    $('#updateForm_' + id).attr('class', '');
    $('#p_comment_' + id).addClass('d-none');
    console.log(id);
}

updateComment = function (id) {
    var originData = $("#updateForm_" + id).serializeArray();
    originData[1] = { name: "commentID", value: id };
    data = JSON.stringify(convertJSON(originData))
    $.ajax({
        type: "put",
        url: "../api/comment",
        data: data,
        headers: {
            'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
        },
        success: function (response) {
            // console.log(response);
            location.reload();
        }
    });
}

deleteComment = function (id) {
    var originData = [{ name: "commentID", value: id }];
    data = JSON.stringify(convertJSON(originData))
    if (window.confirm("確定要刪除？") == true) {
        $.ajax({
            type: "delete",
            url: "../api/comment",
            data: data,
            headers: {
                'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
            },
            success: function (response) {
                // console.log(response);
                location.reload();
            }
        });
    }
}

function convertJSON(originData) {
    data = {}
    $.map(originData, function (n, i) {
        data[n['name']] = n['value']
    });
    return data
}