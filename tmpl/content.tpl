{{define "content"}}
    <div class="row content">
        <div class="col-md-5"></div>
        <div class="col-md-2 no-bg-img">
            <form action="/" method="POST">
                
                <div class="form-group">
                    <label for="cellid">CELL</label>
                    <input type="number" class="form-control" required="true" autofocus="autofocus" name="cellid" id="cellid" placeholder="cellid number" />                                
                </div>
                <div class="form-group">
                    <label for="lacid">LAC</label>
                    <input type="number" class="form-control" required="true" name="lacid" id="lacid" placeholder="lacid number" />
                </div>
                <div class="text-center">
                <button type="submit" class="btn btn-primary">Search</button>            
                </div>            
            </form>
        </div>
        <div class="col-md-5"></div>
    </div>
    {{ if .Search}}
        {{if .Err }}
            <p>{{.Err}}</p>        
        {{else}}
            {{template "result" .}}                            
        {{end}}        
    {{end}}   
{{end}}
