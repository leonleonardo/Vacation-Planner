import {Component, Inject} from '@angular/core';
import {FormControl, FormGroup,} from '@angular/forms';
import {MatDialog, MAT_DIALOG_DATA} from '@angular/material/dialog';
import { map, catchError, tap } from 'rxjs/operators';

import { Trip } from '../trip';
import { TripsService } from '../trips.service';

import { Destination } from '../destination';

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

  constructor(private tripsService: TripsService, public dialog: MatDialog) {}


  destinationResults: Destination = new Destination();

  openDialog(): void {
    this.dialog.open(DestinationResultDialog, {
      data: {results: this.destinationResults},
    });
  }

  add(): void {
    var dest: string = JSON.stringify(this.tripForm.value.dest!);
    var start: string = JSON.stringify(this.tripForm.value.start!);
    var end: string = JSON.stringify(this.tripForm.value.end!);
    start = start.substring(1,11);
    end = end.substring(1,11);
    dest = dest.substring(1,dest.length-1);
    if (!dest) {return;}
    this.tripsService.getTrip({dest, start, end} as Trip).pipe(
      map((response) => {
        this.destinationResults = response;
      })
    )
    .subscribe(response => {
      this.openDialog();
    });
  }

}

@Component({
  selector: 'destination-dialog',
  templateUrl: 'destination-dialog.component.html',
})
export class DestinationResultDialog {
  constructor(@Inject(MAT_DIALOG_DATA) public data: any) {
    // console.log('data passed in is:', this.data);
  }
}