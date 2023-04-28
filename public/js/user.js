$("#login").click(function (e) {  
    var originData = $("#loginForm").serializeArray();
    data = JSON.stringify(convertJSON(originData))
    console.log(data)
    $.ajax({
        type: "post",
        url: "../api/user/login",
        data: data,
        headers: {
            'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
        },
        success: function (response) {
            if (response["error"] == true) {
                $('#errorMsg').remove()
                $('#loginPassword').after(
                    $('<span/>')
                        .addClass('text-danger font-italic')
                        .attr('id', 'errorMsg')
                        .html(response['errorMsg'])
                );
                return
            }
            $(window).attr('location', 'board');
        },
        error: function (xhr, ajaxOptions, thrownError) {
            alert("unknown error");
        }
    });
    e.preventDefault();
});

$('#sign').click(function (e) {
    console.log("click");
    $(window).attr('location', 'register');
    e.preventDefault();
});

$('#signUp').click(function (e) {
    var originData = $("#registerForm").serializeArray();
    data = JSON.stringify(convertJSON(originData))
    console.log(data)
    $.ajax({
        type: "post",
        url: "../api/user/register",
        data: data,
        headers: {
            'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
        },
        success: function (response) {
            if (response["error"] == true) {
                $('#errorMsg').remove()
                $('#registerAccount').after(
                    $('<span/>')
                        .addClass('text-danger font-italic')
                        .attr('id', 'errorMsg')
                        .html(response['errorMsg'])
                );
                return
            }

            alert("註冊成功");
            $(window).attr('location', 'login');
        },
        error: function (xhr, ajaxOptions, thrownError) {
            console.log(xhr);
            console.log(thrownError);
        }
    });
    e.preventDefault();
});

function convertJSON(originData) { 
    data = {}
    $.map(originData, function (n, i) {
        data[n['name']] = n['value']
    });
    return data
}