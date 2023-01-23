import { Injectable } from '@angular/core';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { map } from 'rxjs';

@Injectable( )
export class HelloWorldService {

  constructor(private http: HttpClient) { }

  getTitle() {
    return this.http.get(`${environment.serverUrl}/hello-world`)
    .pipe(map(response => response));
  }
}
