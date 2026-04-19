-- Создаём таблицу организаций (если нет)
CREATE TABLE IF NOT EXISTS organizations (
    id         SERIAL PRIMARY KEY,
    name       TEXT NOT NULL,
    is_default BOOLEAN NOT NULL DEFAULT false
);

-- Добавляем колонку в departments (если нет)
ALTER TABLE departments ADD COLUMN IF NOT EXISTS organization_id INT REFERENCES organizations(id) ON DELETE SET NULL;

BEGIN;

INSERT INTO organizations (name) VALUES ('АДМИНИСТРАЦИИ МО');
INSERT INTO organizations (name) VALUES ('АППАРАТ ГУБЕРНАТОРА И ПРАВИТЕЛЬСТВА ЧАО');
INSERT INTO organizations (name) VALUES ('АРХИВ');
INSERT INTO organizations (name) VALUES ('АТА Чукотки');
INSERT INTO organizations (name) VALUES ('Анадырская транспортная компания');
INSERT INTO organizations (name) VALUES ('АрктикПромПарк');
INSERT INTO organizations (name) VALUES ('АрктикРегионСвязь');
INSERT INTO organizations (name) VALUES ('Библиотека');
INSERT INTO organizations (name) VALUES ('Ветеринарный и фитосанитарный надзор');
INSERT INTO organizations (name) VALUES ('Ветеринарный и фитосанитарный отдел');
INSERT INTO organizations (name) VALUES ('ГБУ Анадырский окружной ПНИ');
INSERT INTO organizations (name) VALUES ('ГБУЗ ЧОБ');
INSERT INTO organizations (name) VALUES ('ГКУ Архив');
INSERT INTO organizations (name) VALUES ('ГКУ Гражданская защита');
INSERT INTO organizations (name) VALUES ('ГКУ Управление кап строительства');
INSERT INTO organizations (name) VALUES ('ГКУ Центр цифр развития');
INSERT INTO organizations (name) VALUES ('ГКУ Центр цифрового развития и ИБ');
INSERT INTO organizations (name) VALUES ('ГКУ ЧАО Чукотуправтодор');
INSERT INTO organizations (name) VALUES ('ГУ МЧС по ЧАО');
INSERT INTO organizations (name) VALUES ('ДЕПАРТАМЕНТЫ ЧАО');
INSERT INTO organizations (name) VALUES ('ДУМА');
INSERT INTO organizations (name) VALUES ('ДШИ');
INSERT INTO organizations (name) VALUES ('Дом народного творчества (ДНТ)');
INSERT INTO organizations (name) VALUES ('ИЗБИРКОМ');
INSERT INTO organizations (name) VALUES ('КОМИТЕТЫ ЧАО');
INSERT INTO organizations (name) VALUES ('Казначество');
INSERT INTO organizations (name) VALUES ('Контакты детские сады');
INSERT INTO organizations (name) VALUES ('Контакты организаций ЧАО');
INSERT INTO organizations (name) VALUES ('Ладушки детский сад');
INSERT INTO organizations (name) VALUES ('Лицей');
INSERT INTO organizations (name) VALUES ('МБОУ школа Сиреники Провиденский ГО');
INSERT INTO organizations (name) VALUES ('МУП СХП, ТСО');
INSERT INTO organizations (name) VALUES ('МЧС');
INSERT INTO organizations (name) VALUES ('Музейный центр');
INSERT INTO organizations (name) VALUES ('НО Фонд развития Чукотки');
INSERT INTO organizations (name) VALUES ('Национальная Гвардия');
INSERT INTO organizations (name) VALUES ('ОДЮШ');
INSERT INTO organizations (name) VALUES ('ООО Система');
INSERT INTO organizations (name) VALUES ('ОСФР ЧАО');
INSERT INTO organizations (name) VALUES ('Офисы Налоговой');
INSERT INTO organizations (name) VALUES ('ПК Полярный');
INSERT INTO organizations (name) VALUES ('ПРОКУРАТУРА');
INSERT INTO organizations (name) VALUES ('РОО ПРЕДПРИНИМАТЕЛИ ЧАО');
INSERT INTO organizations (name) VALUES ('РОСПРИРОДНАДЗОР');
INSERT INTO organizations (name) VALUES ('Росморпорт');
INSERT INTO organizations (name) VALUES ('Роспотребнадзор');
INSERT INTO organizations (name) VALUES ('Росрыболовство');
INSERT INTO organizations (name) VALUES ('СОВЕТ ДЕПУТАТОВ МО');
INSERT INTO organizations (name) VALUES ('СОГАЗ-МЕД');
INSERT INTO organizations (name) VALUES ('Сбербанк');
INSERT INTO organizations (name) VALUES ('Следственное управление');
INSERT INTO organizations (name) VALUES ('Служба судебных приставов');
INSERT INTO organizations (name) VALUES ('Суд ЧАО');
INSERT INTO organizations (name) VALUES ('Судебный департамент');
INSERT INTO organizations (name) VALUES ('Счетная палата');
INSERT INTO organizations (name) VALUES ('Территориальный отдел водных ресурсов по ЧАО');
INSERT INTO organizations (name) VALUES ('УМВД');
INSERT INTO organizations (name) VALUES ('УФАС');
INSERT INTO organizations (name) VALUES ('УФК (Казначейство)');
INSERT INTO organizations (name) VALUES ('УФНС');
INSERT INTO organizations (name) VALUES ('Управление Судебного департамента в ЧАО');
INSERT INTO organizations (name) VALUES ('Управление капитального строительства');
INSERT INTO organizations (name) VALUES ('Управление мировых судей');
INSERT INTO organizations (name) VALUES ('Управление молодежной политики');
INSERT INTO organizations (name) VALUES ('Участки мировых судей');
INSERT INTO organizations (name) VALUES ('ФГУП Почта России');
INSERT INTO organizations (name) VALUES ('ФКП АЭРОПОРТ');
INSERT INTO organizations (name) VALUES ('ФОМС');
INSERT INTO organizations (name) VALUES ('ФСС');
INSERT INTO organizations (name) VALUES ('Фонд развития Чукотки');
INSERT INTO organizations (name) VALUES ('Центр гигиены и эпид');
INSERT INTO organizations (name) VALUES ('ЧМК');
INSERT INTO organizations (name) VALUES ('ЧУКОТКОММУНХОЗ');
INSERT INTO organizations (name) VALUES ('ЧУКОТСНАБ');
INSERT INTO organizations (name) VALUES ('ЧУКОТУПРАВТОДОР');
INSERT INTO organizations (name) VALUES ('ЧФ СВФУ');
INSERT INTO organizations (name) VALUES ('ЧукотАВИА');
INSERT INTO organizations (name) VALUES ('ЧукотЖилосервис');
INSERT INTO organizations (name) VALUES ('Чукотнедра');
INSERT INTO organizations (name) VALUES ('Чукотский институт развития образования');
INSERT INTO organizations (name) VALUES ('Чукотэнерго');
INSERT INTO organizations (name) VALUES ('Эргырон');
INSERT INTO organizations (name) VALUES ('Юстиция');

COMMIT;
