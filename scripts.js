document.querySelectorAll('[data-lang]').forEach(button => {
  button.addEventListener('click', function(e) {
    e.preventDefault();
    
    // Пути к файлам (замените на реальные)
    const filePaths = {
      en: 'path/to/resume_en.pdf',
      ru: 'path/to/resume_ru.pdf'
    };
    
    const lang = this.dataset.lang;
    const fileName = lang === 'en' ? 'resume_en.pdf' : 'resume_ru.pdf';
    
    // Создание скрытой ссылки для скачивания
    const link = document.createElement('a');
    link.href = filePaths[lang];
    link.download = fileName;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  });
});