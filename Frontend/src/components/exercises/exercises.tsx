import ExerciseCard from "@/components/exerciseCard/exerciseCard";
import { exercises as exercisesData } from "@/stubs/exerciseCardStubs";
import { ExerciseCard as ExerciseCardType } from "@/interfaces/interfaces";
import { useEffect, useState } from "react";

// async function GetData(): ExerciseCardType[] {
//   // await new Promise((resolve) => setTimeout(resolve, 1000));
//   let data;
//   try {
//     data = fetch("http://localhost:8080/exercises");
//   } catch (error) {
//     return [];
//   }
//   let exercises = (await data).json();
//   return await exercises;
// }

export default function Exercises() {
  const [data, setData] = useState((): ExerciseCardType[] => []);
  const [isLoading, setLoading] = useState(true);

  useEffect(() => {
    const apiUrl = process.env.NEXT_PUBLIC_API_URL;
    fetch(`${apiUrl}/exercises`)
      .then((res) => res.json())
      .then((data) => {
        setData(data);
        setLoading(false);
      });
  }, []);

  if (isLoading) return <>Loading ...</>;

  return (
    <>
      {data.map((e) => (
        <ExerciseCard key={e.id} exercise={e}></ExerciseCard>
      ))}
    </>
  );
}
