import {Component} from '@angular/core';
import {FormGroup, FormControl} from '@angular/forms';

/** @title Date range picker forms integration */
@Component({
  selector: 'date-range-picker',
  styleUrls: ['date-range-picker.scss'],
  templateUrl: 'date-range-picker.html',
})
export class DateRangePickerComponent {
  range = new FormGroup({
    start: new FormControl<Date | null>(null),
    end: new FormControl<Date | null>(null),
  });
}