package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Content struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price,omitempty"`
	Calories    int     `json:"calories,omitempty"`
}

var dishes = []Content{
	{
		ID:          1,
		Title:       "Surströmming",
		Author:      "Швеция",
		Description: "Шведский национальный деликатес XVI века. Появился во время войны Швеции и Любека: из-за нехватки соли сельдь закисала в бочках. Бедняки попробовали кислую рыбу и оценили её пикантный вкус. Сегодня это квашеная балтийская сельдь с резким специфическим ароматом.",
		Image:       "/assets/images/sur.jpg",
		Price:       1500,
		Calories:    220,
	},
	{
		ID:          2,
		Title:       "Касу-марцу",
		Author:      "Сардиния",
		Description: "Традиционный сардинский сыр с тысячелетней историей. Пастухи Сардинии оставляли овечий сыр пекорино на улице, позволяя сырным мухам откладывать личинки. Ферменты личинок ферментируют жиры, делая текстуру сыра невероятно мягкой и кремообразной.",
		Image:       "/assets/images/kasu-martsu.webp",
		Price:       2100,
		Calories:    380,
	},
	{
		ID:          3,
		Title:       "Дуриан",
		Author:      "Юго-Восточная Азия",
		Description: "«Король фруктов» из Юго-Восточной Азии, известный европейцам со времен путешественников XV века. Известен контрастом: отталкивающий въедливый запах компенсируется божественным нежным вкусом, напоминающим заварной крем с миндалем.",
		Image:       "/assets/images/dur.jpg",
		Price:       950,
		Calories:    147,
	},
	{
		ID:          4,
		Title:       "Хаукарль",
		Author:      "Исландия",
		Description: "Суровое блюдо исландских викингов. Мясо гренландской акулы ядовито в свежем виде из-за высокого содержания мочевины. Чтобы выжить в Арктике, викинги научились ферментировать мясо в гравийных ямах в течение 3–6 месяцев, после чего вялить его на ветру.",
		Image:       "/assets/images/hakarl.jpg",
		Price:       1800,
		Calories:    130,
	},
	{
		ID:          5,
		Title:       "Столетнее яйцо",
		Author:      "Китай",
		Description: "Древнекитайская закуска эпохи династии Мин (около 600 лет назад). Согласно легенде, крестьянин нашел утиные яйца в луже с гашеной известью и чаем, попробовал их и усовершенствовал рецепт. Яйца выдерживают в щелочной смеси из чая, соли и золы несколько месяцев.",
		Image:       "/assets/images/century_egg.jpg",
		Price:       600,
		Calories:    185,
	},
	{
		ID:          6,
		Title:       "Рыба Фугу",
		Author:      "Япония",
		Description: "Японский деликатес с 2300-летней историей. В период Эдо блюдо запрещали из-за частых отравлений самураев. Сегодня его готовят только лицензированные шефы, филигранно удаляющие смертельный тетродотоксин из внутренних органов.",
		Image:       "/assets/images/fugu.jpg",
		Price:       5000,
		Calories:    110,
	},
	{
		ID:          7,
		Title:       "Саннакчи",
		Author:      "Корея",
		Description: "Традиционное корейское блюдо, восходящее к эпохе Корё. Свежего осьминога нарезают непосредственно перед подачей и сбрызгивают кунжутным маслом. Исторически считалось источником выносливости, мужской силы и крепкого здоровья.",
		Image:       "/assets/images/sanakchi.jpg",
		Price:       1400,
		Calories:    90,
	},
	{
		ID:          8,
		Title:       "Вонючий тофу",
		Author:      "Китай",
		Description: "Легенда династии Цин гласит, что торговец Ван Чжихэ забыл нарезанный тофу в кувшине со специями. Тофу забродил и позеленел, но оказался невероятно вкусным. Сегодня это культовый азиатский стритфуд с глубоким пикантным вкусом.",
		Image:       "/assets/images/tofu.jpg",
		Price:       450,
		Calories:    160,
	},
	{
		ID:          9,
		Title:       "Строганина",
		Author:      "Россия",
		Description: "Традиционное блюдо коренных народов Севера России (якутов, ненцев). Зародилось как способ выживания и защиты от цинги в жестком арктическом климате. Свежезамороженная рыбы строгается тонкими ломтиками и подается с солью и перцем.",
		Image:       "/assets/images/stroganina.jpg",
		Price:       1200,
		Calories:    210,
	},
	{
		ID:          10,
		Title:       "Лютефиск",
		Author:      "Скандинавия",
		Description: "Скандинавский зимний деликатес. По легенде, склад сушеной трески викингов сгорел, а зола смешалась с дождевой водой, образовав щелок. Рыбу вымочили от золы и сварили — так появилась треска с уникальной желеобразной текстурой.",
		Image:       "/assets/images/lut.jpg",
		Price:       1650,
		Calories:    115,
	},
	{
		ID:          11,
		Title:       "Жареный тарантул",
		Author:      "Камбоджа",
		Description: "Популярная закуска из камбоджийского городка Скуон. Массово употреблять пауков начали в 1970-х годах во время голода при режиме Красных Кхмеров. Со временем жареный птицеед с чесноком и перцем стал знаменитым уличным деликатесом.",
		Image:       "/assets/images/tar.jpg",
		Price:       800,
		Calories:    240,
	},
	{
		ID:          12,
		Title:       "Эскамолес",
		Author:      "Мексика",
		Description: "Знаменитая «мексиканская икра». Наследие цивилизации ацтеков, которые собирали личинки гигантских черных муравьев в корнях агавы. Считалось пищей императоров и подается обжаренным в масле с эписотом и тортильями.",
		Image:       "/assets/images/ant.jpg",
		Price:       2300,
		Calories:    270,
	},
}

func main() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Обработка /content и /content/:id
	http.HandleFunc("/content", contentHandler)
	http.HandleFunc("/content/", contentHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Сервер запущен на порту " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func contentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	
	path := strings.TrimPrefix(r.URL.Path, "/content")
	path = strings.Trim(path, "/")

	
	if path == "" {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"data": dishes,
		})
		return
	}

	
	id, err := strconv.Atoi(path)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	for _, item := range dishes {
		if item.ID == id {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"data": item,
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Content not found"})
}