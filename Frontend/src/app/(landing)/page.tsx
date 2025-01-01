import style from "./page.module.css";
export default function Landing() {
  return (
    <div className={style.landing}>
      <div className={style.header}>
        <h1>CALISTHETICS</h1>
      </div>
      <div className={style.postHeader}>
        <p>
          Calisthetics — лучший сервис для поиска физических упражнения и
          тренировочных программ.
        </p>
      </div>
      <div className={style.main}>
        <div className={style.exercises}>
          <img
            className={style.exerciseImage}
            src="/pictures/landingExercise.png"
            alt=""
          />
          <div className={style.exerciseText}>
            <h2>Упражнения</h2>
            <p>
              Сервис предоставляет доступ к каталогу различных физических
              упражнений: начиная от силовых и статических со своим весом,
              заканчивая упражнениями на развитие растяжки и баланса.
            </p>
            <a className={style.link} href="/exercises">
              Перейти к упражнениям &rsaquo;
            </a>
          </div>
        </div>

        <div className={style.programs}>
          <div className={style.programText}>
            <h2>Программы</h2>
            <p>
              Сервис предоставляет доступ к каталогу различных тренировочных
              программ с подробной информацией о выполнении упражнений,
              количеством подходов и временем перерыва между ними.
            </p>
            <a className={style.link} href="/programs">
              Перейти к программам &rsaquo;
            </a>
          </div>
          <img
            className={style.programImage}
            src="/pictures/landingProgram.png"
            alt=""
          />
        </div>
      </div>
      <hr />
      <div className={style.description}>
        <p>
          Веб-приложение было сделано в рамках курсового проекта по дисциплине
          "Разработка веб-приложений" студентом 3 курса. В ходе выполнения
          работы использовалась библиотека React, фреймворк Next.JS, NoSQL база
          данных MongoDB, а также язык программирования Go в связке с
          фреймворком Gin. Целью разработки было создание удобного и
          современного сервиса для поиска физических упражнений и программ
          тренировок. Клиентская часть обеспечивает быструю загрузку и
          маршрутизацию благодаря Next.js, а серверная часть реализована на Go
          для достижения высокой производительности. Данные хранятся в MongoDB,
          что позволило гибко структурировать информацию. Проект демонстрирует
          навыки разработки полного стека и может быть расширен в будущем за
          счет новых функций, таких как аналитика и персонализированные
          программы тренировок.
        </p>
      </div>
    </div>
  );
}
