import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { MaterialModule} from './material/material.module';

import { AppComponent } from './app.component';
import { DestinationList } from './components/destination-list';
import { DestinationAdd } from './components/destination-add';



@NgModule({
  declarations: [
    AppComponent,
    DestinationAdd,
    DestinationList 
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    AppRoutingModule,
    MaterialModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
