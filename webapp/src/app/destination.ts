export class Destination {
  Location: string[] = [];
  Restuarants: Description[] = [];
  Enterntainment: Description[] = [];
  Shopping: Description[] = [];
}

export class Description {
  Name: string = '';
  Price: string = '';
  Rating: number = -1;
  Location: string = '';
  Type: string = '';
}


