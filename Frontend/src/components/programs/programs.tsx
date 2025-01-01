import { ProgramsContext } from "@/app/(catalog)/programs/page";
import Card from "@/components/programCard/programCard";
import { ProgramCard } from "@/interfaces/interfaces";
import { useContext, useEffect, useState } from "react";

export default function Programs() {
  const [data, setData] = useState((): ProgramCard[] => []);
  const [isLoading, setLoading] = useState(true);
  const paramContext = useContext(ProgramsContext);

  useEffect(() => {
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;
    fetch(`http://${apiUrl}/programs${paramContext}`)
      .then((res) => res.json())
      .then((data) => {
        setData(data);
        setLoading(false);
      });
  }, [paramContext]);

  if (isLoading) return <>Loading ...</>;
  if (data.length == 0) return <>Таких тренировочных программ нет</>;

  return (
    <>
      {data.map((p) => (
        <Card key={p.id} Program={p}></Card>
      ))}
    </>
  );
}
