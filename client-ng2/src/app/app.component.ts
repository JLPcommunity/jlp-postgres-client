import { Component } from '@angular/core';
import { Http } from '@angular/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  data = [];

  constructor(private http: Http) {
    this.getData();
  }

  getData() {
    this.http.get('http://localhost:8383/get').subscribe(response => {
      console.log(response);
      this.data = response.json();
    }, error => {
      console.log(error);
    });
  }
}
