<!DOCTYPE html>
<html ng-app="ausAnimalApp">
  <head>
    <meta charset="utf-8">
    
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/css/bootstrap.min.css">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-select/1.11.2/css/bootstrap-select.min.css">
    <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.4/js/bootstrap.min.js"></script>
    <!--JavaScript -->
    <link rel="stylesheet" href="static/css/style.css">

    <title>Chennan</title>
  </head>
  <body>
    <header class="container">
      <h1>Chennan Gu Assignment</h1>
    </header>
    <hr>
    <section>
      <div class="container">
        {{range $key, $val := .login}}
          
          <h4>{{$key}}</h4>
          <p>{{$val}}</p>
        {{end}} 
      </div>
    </section>
    <hr>
    <section>
      <div class="container">
        {{range $key, $val := .user}}
          <h4>{{$key}}</h4>
          <p>{{$val}}</p>
        {{end}} 
      </div>
    </section>
    <hr>
    <section>
      <div class="container">
        <h4>Type</h4>
        <p>{{.session.Type}}</p>

        <h4>Content</h4>
        <div class="col-md-6">
        {{range $key, $elem := (.session).Content}}
          <h5>ID</h5>
          <p>{{$elem.ID}}</p>

          <h5>UserId</h5>
          <p>{{$elem.UserId}}</p>

          <h5>CaregiverID</h5>
          <p>{{$elem.CaregiverID}}</p>

          <h5>Created</h5>
          <p>{{$elem.Created}}</p>

          <h5>Latitude</h5>
          <p>{{$elem.Latitude}}</p>

          <h5>Longitud</h5>
          <p>{{$elem.Longitud}}</p>

          <h5>Location</h5>
          <p>{{$elem.Location}}</p>

          <h5>Note</h5>
          <p>{{$elem.Note}}</p>

          <p>...</p>
        {{end}} 
        </div>
      </div>

    </section>
  </body>
</html>