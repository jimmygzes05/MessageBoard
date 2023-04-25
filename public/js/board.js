$('#logout').click(function (e) {
    $(window).attr('location', 'login');
    e.preventDefault();
});

$('#postButton').click(function (e) {
    var data = $("#commentForm").serializeArray();
    userData = { name: "userName", value: sessionStorage.getItem('userName') };
    data[1] = userData;
    console.log(data);
    $.ajax({
        type: "post",
        url: "comment",
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
    var data = $("#updateForm_" + id).serializeArray();
    data[1] = { name: "id", value: id };
    console.log(data);
    $.ajax({
        type: "post",
        url: "updateComment",
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
    var data = [{ name: "id", value: id }];
    if (window.confirm("確定要刪除？") == true) {
        $.ajax({
            type: "post",
            url: "deleteComment",
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