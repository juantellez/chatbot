# chatbot


bot_a.go: Contiene la implementación del bot A.

bot_b.go: Contiene la implementación del bot B.

interaction_service.go: Contiene la implementación del servicio de interacción.

value_provider.go: Contiene la implementación del módulo ValueProvider que proporciona valor al usuario.

payment_processor.go: Contiene la implementación del módulo PaymentProcessor que se encarga de procesar los pagos con Bitcoin y Lightning Network utilizando la API de LNbits.

database.go: Contiene la implementación de las funciones de conexión a la base de datos y las operaciones relacionadas con el almacenamiento de datos.

config.go: Contiene la implementación del módulo Config que carga y maneja la configuración del sistema.

main.go: Es el archivo principal que inicia y ejecuta los servicios y la lógica principal del bot.

config.json: Archivo de configuración que contiene los datos necesarios para la configuración del bot y la conexión a la base de datos.

email_templates/: Directorio que contiene las plantillas de correo electrónico para la verificación de identidad.

web/: Directorio que contiene los archivos HTML y CSS para la interfaz web del cliente.