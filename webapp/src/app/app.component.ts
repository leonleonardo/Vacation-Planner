import { Component, OnInit } from '@angular/core';
import { LocationStrategy } from '@angular/common';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit{
    title = 'Vacation Planner';
    constructor(
      private url:LocationStrategy,
    ) {}

    ngOnInit() {
      console.log(this.url.path());
    }
    
    isRoute(route: string) {
      return (this.url.path()===(route));
    }
  }



