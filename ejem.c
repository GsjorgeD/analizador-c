1#include <stdio.h>


	main() /* Muestra un menú si no se pulsa 4 */
	{
		 char seleccion
		 int var;

     d{
	      printf("1.- Comenzar\n");
	      printf("2.- Abrir\n");
	      printf("3.- Grabar\n");
	      printf("4.- Salir\n");
	      printf("Escoge una opción: ");
	      seleccion=getchar();
	      switch(seleccion){
		 case '1':printf("Opción 1");
		    break;
		 case '2':printf("Opción 2");
		    break;
		 case '3':printf("Opción 3");
		 }

   }while (seleccion!='4') ;

	}
