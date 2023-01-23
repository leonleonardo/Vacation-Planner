import { Component, OnInit } from '@angular/core';
import { HelloWorldService } from './hello-world.service';
import { map } from 'rxjs';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit {
  title = 'ProjFolder';

  constructor (private hw: HelloWorldService) {}

  ngOnInit(): void {
      this.hw.getTitle()
      .pipe(map(data => this.title));

      console.log(this.title);
  }
}
