import {Component} from '@angular/core';
import {FormControl, FormGroup,} from '@angular/forms';

import { Trip } from '../trip';
import { TripsService } from '../trips.service';

/** @title Date range picker forms integration */
@Component({
  selector: 'trip-search',
  styleUrls: ['trip-search.component.scss'],
  templateUrl: 'trip-search.component.html',
})
export class TripSearchComponent {

  public tripForm = new FormGroup({
    dest: new FormControl<string | null>(null),
    start: new FormControl<string | null>(null),
    end: new FormControl<string | null>(null),
  });

  constructor(private tripsService: TripsService) {}

  destInfo = new Array<any>();

  add(): void {
    var dest: string = JSON.stringify(this.tripForm.value.dest!);
    var start: string = JSON.stringify(this.tripForm.value.start!);
    var end: string = JSON.stringify(this.tripForm.value.end!);
    start = start.substring(1,11);
    end = end.substring(1,11);
    dest = dest.substring(1,dest.length-1);
    // console.log(start);
    // console.log(end);
    // console.log(dest);
    if (!dest) {return;}
    this.tripsService.getTrip({dest, start, end} as Trip)
    .subscribe(response => {
      this.destInfo = response.data;
      console.log(this.destInfo);
    });
  }

}