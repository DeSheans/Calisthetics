import Card from "@/components/programCard/programCard";
import { ProgramCard } from "@/interfaces/interfaces";
import { twoDaySplit } from "@/stubs/programStub";

function GetData(): ProgramCard[] {
  // await new Promise((resolve) => setTimeout(resolve, 3000));
  return [twoDaySplit];
}

export default function Programs() {
  let programs = GetData();

  return (
    <>
      {programs.map((p) => (
        <Card key={p.id} Program={p}></Card>
      ))}
    </>
  );
}
