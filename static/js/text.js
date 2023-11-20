function deleteTextRecord() {
    let items=document.getElementsByClassName('cb');
    let len=items.length;
    for (var i=len-1; i>=0;i--) {
        let is_checkd = items[i].checked;
        if (is_checkd) {
            let divItems = items[i].parentNode.parentNode;
            let url = divItems.innerText;
            $.ajax(
                {
                    url: "/svc/text/record/del",
                    type: 'POST',
                    // contentType: 'application/json',
                    data: {
                        "url": url,
                    }
                },

            );
            divItems.parentNode.removeChild(divItems);
        }
    }
}


function checkEsConn() {
    $.ajax({
        url: '/svc/text/es/ck',
        type: 'GET',
        success: function(data) {
            console.log('Response:', data);
            if (data["code"] !== 200){
                alert(data["msg"])
                return
            }
            alert(data["msg"])
        },
        error: function(error) {
            console.error('Error:', error);
            return
        }
    });
}