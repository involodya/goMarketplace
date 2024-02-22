import {useEffect, useState} from "react";
import {Link} from "react-router-dom";
import {NoneProductImage} from "./icons/NoneProductImage";
import {API_URL} from './config';
import Hat from './Hat';
import './MainPage.css';


const ProductCard = ({productImageUrl, productId, productName, price}) => {
    return (
        <Link to={`/product/${productId}`}>
            <div className="cards-group">
                <div className="product-card">
                    <div className="name-wrapper">
                        <div className="product-name">{productName}</div>
                    </div>
                    <img src={productImageUrl}
                         alt={productName} className="product-image"/>
                    {/*<NoneProductImage className="vuesax-bold-gallery"/>*/}
                    <div className="product-price">${price}</div>
                </div>
            </div>
        </Link>
    );
};

function MainPage() {
    const [products, setProducts] = useState([{ID: 1, Name: "", Cost: 0}]);
    const [userStats, setUserStats] = useState({item_count: 0});

    useEffect(() => {
        (async () => {
            const token = window.localStorage.getItem("access");
            const response = await fetch(`${API_URL}/items`, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });
            const data = await response.json();
            setProducts(data);
        })();
        (async () => {
            const token = window.localStorage.getItem("access");
            const response = await fetch(`${API_URL}/stats`, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });
            const data = await response.json();
            setUserStats(data);
        })();
    }, []);

    const renderProductCards = () => (
        <div className="products">
            {products.map((product, index) => (
                <ProductCard
                    key={index}
                    productId={product.ID}
                    productImageUrl={product.ImageUrl}
                    productName={product.Name}
                    price={product.Cost}
                />
            ))}
        </div>
    );

    const renderCard = () => (
        <section className="content">
            <div className="label">Statistics</div>
            <div className="statistics">
                <div className="text-wrapper">Total items sold: {userStats.item_count}</div>
            </div>
            <div className="label">Your Products</div>
            {renderProductCards()}
        </section>
    );

    return (
        <main className="main-page">
            <Hat/>
            {renderCard()}
        </main>
    );
}

export default MainPage;