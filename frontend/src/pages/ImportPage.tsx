import { useState } from "react";
import { Button, Input, Container } from "semantic-ui-react";

export const ImportPage: React.FC = () => {
    const [file, setFile] = useState<File | null>(null);

    const handleUpload = async () => {
        if (!file) {
            alert("Выберите файл для загрузки");
            return;
        }

        const formData = new FormData();
        formData.append("excel", file);

        try {
            const response = await fetch("/upload", {
                method: "POST",
                body: formData,
            });

            if (!response.ok) {
                const errorData = await response.json();
                alert("Ошибка при загрузке: " + (errorData.error || response.statusText));
                return;
            }

            alert("Файл успешно загружен!");
        } catch (err) {
            console.error(err);
            alert("Ошибка сети или сервера");
        }
    };

    return (
        <Container>
            <h2>Upload Excel</h2>
            <Input type="file" onChange={(e) => setFile(e.target.files?.[0] || null)} />
            <Button onClick={handleUpload} primary>Upload</Button>
        </Container>
    );
};
