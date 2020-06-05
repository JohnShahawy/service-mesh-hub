// Code generated by skv2. DO NOT EDIT.

package v1alpha1sets

import (
	. "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"

	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

type KubernetesClusterSet interface {
	Keys() sets.String
	List() []*KubernetesCluster
	Map() map[string]*KubernetesCluster
	Insert(kubernetesCluster ...*KubernetesCluster)
	Equal(kubernetesClusterSet KubernetesClusterSet) bool
	Has(kubernetesCluster *KubernetesCluster) bool
	Delete(kubernetesCluster *KubernetesCluster)
	Union(set KubernetesClusterSet) KubernetesClusterSet
	Difference(set KubernetesClusterSet) KubernetesClusterSet
	Intersection(set KubernetesClusterSet) KubernetesClusterSet
}

func makeGenericKubernetesClusterSet(kubernetesClusterList []*KubernetesCluster) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range kubernetesClusterList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type kubernetesClusterSet struct {
	set sksets.ResourceSet
}

func NewKubernetesClusterSet(kubernetesClusterList ...*KubernetesCluster) KubernetesClusterSet {
	return &kubernetesClusterSet{set: makeGenericKubernetesClusterSet(kubernetesClusterList)}
}

func (s kubernetesClusterSet) Keys() sets.String {
	return s.set.Keys()
}

func (s kubernetesClusterSet) List() []*KubernetesCluster {
	var kubernetesClusterList []*KubernetesCluster
	for _, obj := range s.set.List() {
		kubernetesClusterList = append(kubernetesClusterList, obj.(*KubernetesCluster))
	}
	return kubernetesClusterList
}

func (s kubernetesClusterSet) Map() map[string]*KubernetesCluster {
	newMap := map[string]*KubernetesCluster{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*KubernetesCluster)
	}
	return newMap
}

func (s kubernetesClusterSet) Insert(
	kubernetesClusterList ...*KubernetesCluster,
) {
	for _, obj := range kubernetesClusterList {
		s.set.Insert(obj)
	}
}

func (s kubernetesClusterSet) Has(kubernetesCluster *KubernetesCluster) bool {
	return s.set.Has(kubernetesCluster)
}

func (s kubernetesClusterSet) Equal(
	kubernetesClusterSet KubernetesClusterSet,
) bool {
	return s.set.Equal(makeGenericKubernetesClusterSet(kubernetesClusterSet.List()))
}

func (s kubernetesClusterSet) Delete(KubernetesCluster *KubernetesCluster) {
	s.set.Delete(KubernetesCluster)
}

func (s kubernetesClusterSet) Union(set KubernetesClusterSet) KubernetesClusterSet {
	return NewKubernetesClusterSet(append(s.List(), set.List()...)...)
}

func (s kubernetesClusterSet) Difference(set KubernetesClusterSet) KubernetesClusterSet {
	newSet := s.set.Difference(makeGenericKubernetesClusterSet(set.List()))
	return kubernetesClusterSet{set: newSet}
}

func (s kubernetesClusterSet) Intersection(set KubernetesClusterSet) KubernetesClusterSet {
	newSet := s.set.Intersection(makeGenericKubernetesClusterSet(set.List()))
	var kubernetesClusterList []*KubernetesCluster
	for _, obj := range newSet.List() {
		kubernetesClusterList = append(kubernetesClusterList, obj.(*KubernetesCluster))
	}
	return NewKubernetesClusterSet(kubernetesClusterList...)
}

type MeshServiceSet interface {
	Keys() sets.String
	List() []*MeshService
	Map() map[string]*MeshService
	Insert(meshService ...*MeshService)
	Equal(meshServiceSet MeshServiceSet) bool
	Has(meshService *MeshService) bool
	Delete(meshService *MeshService)
	Union(set MeshServiceSet) MeshServiceSet
	Difference(set MeshServiceSet) MeshServiceSet
	Intersection(set MeshServiceSet) MeshServiceSet
}

func makeGenericMeshServiceSet(meshServiceList []*MeshService) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range meshServiceList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type meshServiceSet struct {
	set sksets.ResourceSet
}

func NewMeshServiceSet(meshServiceList ...*MeshService) MeshServiceSet {
	return &meshServiceSet{set: makeGenericMeshServiceSet(meshServiceList)}
}

func (s meshServiceSet) Keys() sets.String {
	return s.set.Keys()
}

func (s meshServiceSet) List() []*MeshService {
	var meshServiceList []*MeshService
	for _, obj := range s.set.List() {
		meshServiceList = append(meshServiceList, obj.(*MeshService))
	}
	return meshServiceList
}

func (s meshServiceSet) Map() map[string]*MeshService {
	newMap := map[string]*MeshService{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*MeshService)
	}
	return newMap
}

func (s meshServiceSet) Insert(
	meshServiceList ...*MeshService,
) {
	for _, obj := range meshServiceList {
		s.set.Insert(obj)
	}
}

func (s meshServiceSet) Has(meshService *MeshService) bool {
	return s.set.Has(meshService)
}

func (s meshServiceSet) Equal(
	meshServiceSet MeshServiceSet,
) bool {
	return s.set.Equal(makeGenericMeshServiceSet(meshServiceSet.List()))
}

func (s meshServiceSet) Delete(MeshService *MeshService) {
	s.set.Delete(MeshService)
}

func (s meshServiceSet) Union(set MeshServiceSet) MeshServiceSet {
	return NewMeshServiceSet(append(s.List(), set.List()...)...)
}

func (s meshServiceSet) Difference(set MeshServiceSet) MeshServiceSet {
	newSet := s.set.Difference(makeGenericMeshServiceSet(set.List()))
	return meshServiceSet{set: newSet}
}

func (s meshServiceSet) Intersection(set MeshServiceSet) MeshServiceSet {
	newSet := s.set.Intersection(makeGenericMeshServiceSet(set.List()))
	var meshServiceList []*MeshService
	for _, obj := range newSet.List() {
		meshServiceList = append(meshServiceList, obj.(*MeshService))
	}
	return NewMeshServiceSet(meshServiceList...)
}

type MeshWorkloadSet interface {
	Keys() sets.String
	List() []*MeshWorkload
	Map() map[string]*MeshWorkload
	Insert(meshWorkload ...*MeshWorkload)
	Equal(meshWorkloadSet MeshWorkloadSet) bool
	Has(meshWorkload *MeshWorkload) bool
	Delete(meshWorkload *MeshWorkload)
	Union(set MeshWorkloadSet) MeshWorkloadSet
	Difference(set MeshWorkloadSet) MeshWorkloadSet
	Intersection(set MeshWorkloadSet) MeshWorkloadSet
}

func makeGenericMeshWorkloadSet(meshWorkloadList []*MeshWorkload) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range meshWorkloadList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type meshWorkloadSet struct {
	set sksets.ResourceSet
}

func NewMeshWorkloadSet(meshWorkloadList ...*MeshWorkload) MeshWorkloadSet {
	return &meshWorkloadSet{set: makeGenericMeshWorkloadSet(meshWorkloadList)}
}

func (s meshWorkloadSet) Keys() sets.String {
	return s.set.Keys()
}

func (s meshWorkloadSet) List() []*MeshWorkload {
	var meshWorkloadList []*MeshWorkload
	for _, obj := range s.set.List() {
		meshWorkloadList = append(meshWorkloadList, obj.(*MeshWorkload))
	}
	return meshWorkloadList
}

func (s meshWorkloadSet) Map() map[string]*MeshWorkload {
	newMap := map[string]*MeshWorkload{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*MeshWorkload)
	}
	return newMap
}

func (s meshWorkloadSet) Insert(
	meshWorkloadList ...*MeshWorkload,
) {
	for _, obj := range meshWorkloadList {
		s.set.Insert(obj)
	}
}

func (s meshWorkloadSet) Has(meshWorkload *MeshWorkload) bool {
	return s.set.Has(meshWorkload)
}

func (s meshWorkloadSet) Equal(
	meshWorkloadSet MeshWorkloadSet,
) bool {
	return s.set.Equal(makeGenericMeshWorkloadSet(meshWorkloadSet.List()))
}

func (s meshWorkloadSet) Delete(MeshWorkload *MeshWorkload) {
	s.set.Delete(MeshWorkload)
}

func (s meshWorkloadSet) Union(set MeshWorkloadSet) MeshWorkloadSet {
	return NewMeshWorkloadSet(append(s.List(), set.List()...)...)
}

func (s meshWorkloadSet) Difference(set MeshWorkloadSet) MeshWorkloadSet {
	newSet := s.set.Difference(makeGenericMeshWorkloadSet(set.List()))
	return meshWorkloadSet{set: newSet}
}

func (s meshWorkloadSet) Intersection(set MeshWorkloadSet) MeshWorkloadSet {
	newSet := s.set.Intersection(makeGenericMeshWorkloadSet(set.List()))
	var meshWorkloadList []*MeshWorkload
	for _, obj := range newSet.List() {
		meshWorkloadList = append(meshWorkloadList, obj.(*MeshWorkload))
	}
	return NewMeshWorkloadSet(meshWorkloadList...)
}

type MeshSet interface {
	Keys() sets.String
	List() []*Mesh
	Map() map[string]*Mesh
	Insert(mesh ...*Mesh)
	Equal(meshSet MeshSet) bool
	Has(mesh *Mesh) bool
	Delete(mesh *Mesh)
	Union(set MeshSet) MeshSet
	Difference(set MeshSet) MeshSet
	Intersection(set MeshSet) MeshSet
}

func makeGenericMeshSet(meshList []*Mesh) sksets.ResourceSet {
	var genericResources []metav1.Object
	for _, obj := range meshList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type meshSet struct {
	set sksets.ResourceSet
}

func NewMeshSet(meshList ...*Mesh) MeshSet {
	return &meshSet{set: makeGenericMeshSet(meshList)}
}

func (s meshSet) Keys() sets.String {
	return s.set.Keys()
}

func (s meshSet) List() []*Mesh {
	var meshList []*Mesh
	for _, obj := range s.set.List() {
		meshList = append(meshList, obj.(*Mesh))
	}
	return meshList
}

func (s meshSet) Map() map[string]*Mesh {
	newMap := map[string]*Mesh{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*Mesh)
	}
	return newMap
}

func (s meshSet) Insert(
	meshList ...*Mesh,
) {
	for _, obj := range meshList {
		s.set.Insert(obj)
	}
}

func (s meshSet) Has(mesh *Mesh) bool {
	return s.set.Has(mesh)
}

func (s meshSet) Equal(
	meshSet MeshSet,
) bool {
	return s.set.Equal(makeGenericMeshSet(meshSet.List()))
}

func (s meshSet) Delete(Mesh *Mesh) {
	s.set.Delete(Mesh)
}

func (s meshSet) Union(set MeshSet) MeshSet {
	return NewMeshSet(append(s.List(), set.List()...)...)
}

func (s meshSet) Difference(set MeshSet) MeshSet {
	newSet := s.set.Difference(makeGenericMeshSet(set.List()))
	return meshSet{set: newSet}
}

func (s meshSet) Intersection(set MeshSet) MeshSet {
	newSet := s.set.Intersection(makeGenericMeshSet(set.List()))
	var meshList []*Mesh
	for _, obj := range newSet.List() {
		meshList = append(meshList, obj.(*Mesh))
	}
	return NewMeshSet(meshList...)
}
