{{define "base"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">	
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="Cell Search written in Golang">
    <meta name="author" content="Daniel Rittweiler">
    <!-- Bootstrap Core CSS -->
    <link href="/css/bootstrap.min.css" rel="stylesheet">   

    <!-- Custom -->
    <link href="/css/custom.css" rel="stylesheet">
	
    <title>CS|GO</title>
</head>
<body>
    <div class="container-fluid">
        <div class="row">            
            <div id="csgo" class="col-md-offset-4 col-md-4 text-vertical-center">
                <h1>[ CS|GO ]</h1>
                <div><strong>Cell Search written in Go</strong></div>
            </div>            
        </div>        
        {{template "content" .}}        
    </div>      
    <script src="/js/jquery-3.1.1.slim.min.js" type="application/javascript"></script>
    <script src="/js/bootstrap.min.js" type="application/javascript"></script>   
    <script>
        $(document).ready(function () { 
            if (typeof render_map !== "undefined") { 
                render_map();
            }
            
        });
    </script>
    
</body>
</html>
{{end}}
