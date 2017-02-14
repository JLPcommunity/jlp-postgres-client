import { Component, OnInit } from '@angular/core';
import { Http } from '@angular/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  data = [];
  status = '';

  constructor(private http: Http) {
  }

  ngOnInit() {
    this.getData();
  }

  getData() {
    this.status = 'Requesting...';
    this.http.get('http://localhost:8383/get').subscribe(response => {
      console.log(response);
      this.data = response.json();
    }, error => {
      console.log(error);
      // Get the 'cause' from error body to be the status
      var body = error.json()
      if(body != null){
        this.status = body.cause;
      }
    });
  }
}
