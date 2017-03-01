var socket;

$(document).ready(function () {
    // Create a socket
    socket = new WebSocket('ws://' + window.location.host + '/join');
    // Message received on the socket
    socket.onmessage = function (event) {
        var data = JSON.parse(event.data);
        if (data.type == "jobflow"){
           var jobid = data.jobid
           var dur = data.dis
           var appenddata = "<p><b>任务ID：" + data.jobid + "</b></p>"
           var getdata = "<p>"+ data.data +"</p>"
           var version = "v1"
           var dataobj = eval('(' + data.data + ')')
           for (var key in dataobj){
              version = dataobj[key]['Version']
           }
           var durtime = "<p><span class='label label-default'>耗时："+ data.dis + "s</span>&nbsp;&nbsp<span class='label label-primary'>版本号："+ version + "</span></p>"
           $('#logoutput').append(appenddata+getdata+durtime)
        }else if (data.type == "summary") {
           var versiondata = data.versionmap
           var version_output = ""
           for (var key in versiondata){
             version_output+="&nbsp;&nbsp;<span class='label label-info'>版本 "+ key + " / " + versiondata[key] + "次</span>&nbsp;&nbsp;&nbsp;"
           }
           var getdata = "<h4>信息汇总：</h4><p id='summary'><span class='label label-warning'>总时间:"+ data.dis+"s</span></p>"
           $('#logoutput').append(getdata)
           $('#summary').append(version_output)
        }
    };

    // Send messages.
    var postConecnt = function () {
        var content = {
          'jobname':$('#jobname').val(),
          'parallel':$('#parallel').val(),
          'svcname':$('#svcname').val(),
          'solution':$('#solution').val()
        };
        var data = JSON.stringify(content)
        socket.send(data);
        //$('#jobname').val("");
        //$('#parallel').val("");
        //$('#svcname').val("");
        //$('#solution').val("");
    }

    $('#sendbtn').click(function () {
        postConecnt();
    });
    $('#clearbtn').click(function () {
      $('#logoutput').text("")
    });
});
