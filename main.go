package main

import (
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	namespace := "wordpresses"
	selector := "app.kubernetes.io/name=wordpress"
	one := int32(1)
	zero := int32(0)
	connect_k8s()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		wpdeps, _ := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{
			LabelSelector: selector,
		})
		c.HTML(200, "index.tpml", gin.H{
			"deps":  wpdeps.Items,
			"debug": wpdeps.Items[0].Status.Replicas,
		})
	})
	r.GET("/start/:name", func(c *gin.Context) {
		depName := c.Param("name")
		dep, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), depName, metav1.GetOptions{})
		if err == nil {
			dep.Spec.Replicas = &one
			println(dep.Name)
			_, err = clientset.AppsV1().Deployments("wordpresses").Update(context.TODO(), dep, metav1.UpdateOptions{})
			if err == nil {
				c.Redirect(302, "/")
			} else {
				panic(err)
			}
		} else {
			panic(err)
		}
	})
	r.GET("/stop/:name", func(c *gin.Context) {
		depName := c.Param("name")
		dep, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), depName, metav1.GetOptions{})
		if err == nil {
			println(dep.Name)
			dep.Spec.Replicas = &zero
			_, err = clientset.AppsV1().Deployments("wordpresses").Update(context.TODO(), dep, metav1.UpdateOptions{})
			if err == nil {
				c.Redirect(302, "/")
			} else {
				panic(err)
			}
		}
	})
	r.Run("0.0.0.0:4000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
