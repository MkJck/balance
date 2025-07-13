import React, { useState } from "react";                              // Импорт React и хука useState, который управляет состоянием компонента

const CreateTransactionPage: React.FC = () => {                       // Создание функционального компонента с 3 состояниями:
  const [amount, setAmount] = useState("100");                        // amount - сумма транзакции, изначально пустая строка
  const [participants, setParticipants] = useState("1, 2");           // --//--
  const [description, setDescription] = useState("Ужин на двоих");    // --//--
                                                                      // Для каждого состояния пара: значение и функция для его изменения

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();                                               // Предотвращает дефолтную реакцию - перезагрузку страницы
    // Здесь будет отправка данных на сервер (позже). Пока данные выводятся через алерт, в будущем будут отправлены в json-е
    alert(`Создана транзакция:
      Сумма: ${amount}
      Участники: ${participants}
      Описание: ${description}`);
  };

  return (
    <div style={{                                                     // Внешний контейнер с наборами стилей
      maxWidth: 400,
      margin: "2rem auto",                        // Внешние отступы
      padding: "2rem",                            // Внутренние отступы
      borderRadius: "12px",
      background: "#18151bff",
      boxShadow: "0 2px 8px rgba(0,0,0,0.07)"
    }}>
      <h2 style={{ textAlign: "center", marginBottom: "1.5rem", color: "#fff", fontWeight: "bold", fontSize: "24px", }}>Создать транзакцию</h2>
      <form onSubmit={handleSubmit}>
        <div style={{ marginBottom: 12 }}>
          <label>Сумма:</label>
          <input
            type="number"
            value={amount}
            onChange={e => setAmount(e.target.value)}
            required
            style={{ width: "100%", padding: 8, marginTop: 4, borderRadius: "6px", }}
          />
        </div>
        <div style={{ marginBottom: 12 }}>
          <label>Участники (ID через запятую):</label>
          <input
            type="text"
            value={participants}
            onChange={e => setParticipants(e.target.value)}
            placeholder="1,2,3"
            required
            style={{ width: "100%", padding: 8, marginTop: 4, borderRadius: "6px", }}
          />
        </div>
        <div style={{ marginBottom: 12 }}>
          <label>Описание:</label>
          <input
            type="text"
            value={description}
            onChange={e => setDescription(e.target.value)}
            style={{ width: "100%", padding: 8, marginTop: 4, borderRadius: "6px", }}
          />
        </div>
        <button type="submit" style={{
          width: "100%",
          padding: 10,
          background: "#4CAF50",
          color: "#fff",
          border: "none",
          borderRadius: 6,
          fontWeight: "bold",
          cursor: "pointer",
          transition: "background 0.3s",
          fontSize: "16px"
        }}>
          Создать
        </button>
      </form>
    </div>
  );
};

export default CreateTransactionPage;