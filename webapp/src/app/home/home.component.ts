import {Component} from '@angular/core';

import { Trip } from '../trip';
import { TripsService } from '../trips.service';

import {FormControl} from '@angular/forms';

/**
 * @title Basic grid-list
 */
@Component({
  selector: 'app-home',
  styleUrls: ['home.component.scss'],
  templateUrl: 'home.component.html',
})
export class HomeComponent {
  trips: Trip[] = [];

  constructor(private tripsService: TripsService) {}

  date = new FormControl(new Date());

  ngOnInit(): void {
    this.getTrips();
  }

  getTrips(): void {
    this.tripsService.getTrips()
    .subscribe(trips => this.trips = trips);
  }

  add(destination: string, start: string, end: string): void {
    destination = destination.trim();
    start = start.trim();
    end = end.trim();
    if (!destination) {return;}
    this.tripsService.addTrips({destination, start, end} as Trip)
      .subscribe(trip => {
        this.trips.push(trip);
      });
  }

}