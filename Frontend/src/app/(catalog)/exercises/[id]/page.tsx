import { Exercise } from "@/interfaces/interfaces";
import styles from "./page.module.css";
import ExerciseBrief from "@/components/exerciseBrief/exerciseBrief";
import Carousel from "@/components/carousel/carousel";

// const isFavorite = false;
const similar = false;

export default async function ExercisePage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const id = (await params).id;
  const apiUrl = process.env.NEXT_PRIVATE_API_URL;
  const data: Promise<Exercise> = fetch(
    `http://${apiUrl}/exercises/${id}`
  ).then((res) => res.json());
  if (data === null) return <>ERROR 404</>;

  return (
    <div className={styles.inner}>
      <h1 className={styles.heading}>{(await data).name}</h1>

      <div className={styles.carousel}>
        <Carousel
          images={(await data).pictures.map((p) => `/pictures/${p}`)}
        ></Carousel>
      </div>

      <div className={styles.briefContainer}>
        <ExerciseBrief exercise={await data} isFavorite={true}></ExerciseBrief>
      </div>

      <div className={styles.descriptionContainer}>
        <Description
          header="Техника выполнения"
          text={(await data).description}
        ></Description>
        <List header="Советы:" listItems={(await data).tips}></List>
        <List
          header="Варианты усложнения:"
          listItems={(await data).variations}
        ></List>
      </div>

      {similar && (
        <div className={styles.similar}>
          <hr className={styles.hr}></hr>
          <h2>Похожие упражнения</h2>
          <div className={styles.similarItemsContainer}>
            {[0, 0, 0, 0, 0].map(() => (
              <div className={styles.similarItem}></div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}

function Description({ header, text }: { header: string; text: string }) {
  return (
    text && (
      <div className={styles.description}>
        <h2>{header}</h2>
        <p>{text}</p>
      </div>
    )
  );
}

function List({ header, listItems }: { header: string; listItems: string[] }) {
  let c = 0;
  return (
    listItems && (
      <div className={styles.list}>
        <h2>{header}</h2>
        {listItems.map((i) => (
          <li key={c++}>{i}</li>
        ))}
      </div>
    )
  );
}
