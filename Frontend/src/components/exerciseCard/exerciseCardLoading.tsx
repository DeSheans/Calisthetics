import styles from "./exerciseCard.module.css";
export default function ExerciseCardLoading() {
  return (
    <div className={styles.card}>
      <div style={{ height: "21rem" }}></div>
    </div>
  );
}
