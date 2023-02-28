import {Component} from '@angular/core';

@Component({
  selector: 'trips-deck-component',
  templateUrl: 'deck.component.html',
  styleUrls: ['deck.component.scss'],
})
export class TripsCardDeckComponent{
  cardList = [
    {'cardTitle': 'Miami', 'cardSubTitle': '02/16/2023 - 02/18/2023', 'cardContent': 'Content for card 1', 'footer': 'footer for card 1'},
    {'cardTitle': 'Card2 Title', 'cardSubTitle': 'Card2 sub title', 'cardContent': 'Card2 content ', 'footer': 'Card2 footer'},
    {'cardTitle': 'Card3 Title', 'cardSubTitle': 'Card3 sub title', 'cardContent': 'Card3 content', 'footer': 'card3 footer'}

]
}