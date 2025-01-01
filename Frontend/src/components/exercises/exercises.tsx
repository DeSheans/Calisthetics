import { ExercisesContext } from "@/app/(catalog)/exercises/page";
import ExerciseCard from "@/components/exerciseCard/exerciseCard";
import { ExerciseCard as ExerciseCardType } from "@/interfaces/interfaces";
import { useContext, useEffect, useState } from "react";

export default function Exercises() {
  const [data, setData] = useState((): ExerciseCardType[] => []);
  const [isLoading, setLoading] = useState(true);
  const paramContext = useContext(ExercisesContext);

  useEffect(() => {
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;
    fetch(`http://${apiUrl}/exercises${paramContext}`)
      .then((res) => res.json())
      .then((data) => {
        setData(data);
        setLoading(false);
      });
  }, [paramContext]);

  if (isLoading) return <>Loading ...</>;
  if (data.length == 0) return <>Таких упражнений нет</>;

  return (
    <>
      {data.map((e) => (
        <ExerciseCard key={e.id} exercise={e}></ExerciseCard>
      ))}
    </>
  );
}
