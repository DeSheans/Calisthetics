import { Equipment } from "@/interfaces/interfaces";

export const bar: Equipment = {
  id: 0,
  name: "высокая перекладина",
};

export const bars: Equipment = {
  id: 1,
  name: "брусья",
};
export const wallBars: Equipment = {
  id: 2,
  name: "шведская стенка",
};
export const bench: Equipment = {
  id: 3,
  name: "скамья для пресса",
};
export const rings: Equipment = {
  id: 4,
  name: "гимнастические кольца",
};

export const equipment: Equipment[] = [bar, bars, wallBars, bench, rings];
