import ExerciseCardLoading from "../exerciseCard/exerciseCardLoading";

export default function ExercisesLoading() {
  return (
    <>
      {Array(10)
        .fill(0)
        .map(() => (
          <ExerciseCardLoading></ExerciseCardLoading>
        ))}
    </>
  );
}
